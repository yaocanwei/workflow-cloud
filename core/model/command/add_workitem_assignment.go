package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/helper"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

// TypeWorkitemAssignmentAdd 任务指派人定义
const (
	TypeWorkitemAssignmentAdd = "workitem_assignment:add"
)

// NewWorkitemAssignmentAddTask 初始化节点指派人任务
func NewWorkitemAssignmentAddTask(workitemID, partyID int, permanent bool) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
		"party_id":    partyID,
		"permanent":   permanent,
	}
	return cmdx.NewTask(TypeWorkitemAssignmentAdd, payload)
}

// HandleWorkitemAssignmentAddTask 节点指派人
func HandleWorkitemAssignmentAddTask(ctx context.Context, t *cmdx.Task) error {
	var err error
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	partyID, err := t.Payload.GetInt("party_id")
	if err != nil {
		return err
	}
	permanent, err := t.Payload.GetBool("permanent")

	if partyID == 0 {
		return fmt.Errorf("party was required")
	}
	var workitem *model.WfWorkitem
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if permanent {
		t = NewManualAssignmentAddTask(workitem.CaseID, workitem.TransitionID, partyID)
		err = HandleManualAssignmentAddTask(ctx, t)
		return err
	}

	notifiedUsers := make([][]model.WfUser, 0)
	for _, party := range workitem.Parties {
		notifiedUsers = append(notifiedUsers, party.Users)
	}
	// TODO: refactor
	notifiedUsersFlatten := helper.Flatten(notifiedUsers)
	mNotifiedUsers := convToStruct(notifiedUsersFlatten)
	var workitemAssignment *model.WfWorkitemAssignment
	workitemAssignment, err = workitemAssignment.FindByPartyID(workitem.ID, partyID)
	if err == nil {
		return fmt.Errorf("")
	}

	workitemAssignment.PartyID = partyID
	workitemAssignment.WorkitemID = workitem.ID
	workitemAssignment.Create()
	var party *model.WfParty
	party, err = party.FindByID(partyID)
	newUsers := party.Users
	// TODO: refactor
	toNotify := diff(newUsers, mNotifiedUsers)
	transition := workitem.Transition
	fmt.Println("newUsers", newUsers, toNotify, transition)
	for _, notifyUser := range toNotify {
		if transition.MultipleInstance && workitem.Forked {
			if workitem.HasChild(workitemID) {
				continue
			}
			child := &model.WfWorkitem{
				WorkflowID:    workitem.WorkflowID,
				TransitionID:  workitem.TransitionID,
				State:         0,
				TriggerTime:   workitem.TriggerTime,
				Forked:        true,
				HoldingUserID: notifyUser.ID,
				CaseID:        workitem.CaseID,
			}
			child, _ = child.Create()
			// workitem.Transition.NotificationCallback
		} else {

		}
	}
	return nil
}

func convToStruct(data []interface{}) []model.WfUser {
	users := make([]model.WfUser, 0)
	for _, v := range data {
		users = append(users, v.(model.WfUser))
	}
	return users
}

// TODO: 抽象为通用
func diff(newUsers []model.WfUser, notifiedUsers []model.WfUser) []model.WfUser {
	var diff []model.WfUser
	for i := 0; i < 2; i++ {
		for _, newUser := range newUsers {
			found := false
			for _, notifiedUser := range notifiedUsers {
				if newUser.Account == notifiedUser.Account {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, newUser)
			}
		}
		if i == 0 {
			newUsers, notifiedUsers = notifiedUsers, newUsers
		}
	}

	return diff
}
