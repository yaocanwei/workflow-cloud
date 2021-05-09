package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeWorkitemAssignmentRemove = "workitem_assignment:remove"
)

func NewWorkitemAssignmentTask(workitemID, partyID int, permanent bool) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
		"party_id":    partyID,
		"permanent":   permanent,
	}
	return cmdx.NewTask(TypeWorkitemAssignmentRemove, payload)
}

func HandleWorkitemAssignmentTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	partyID, err := t.Payload.GetInt("party_id")
	if err != nil {
		return err
	}
	permanent, err := t.Payload.GetBool("permanent")
	if err != nil {
		return err
	}
	fmt.Println(workitemID, partyID, permanent)
	//var party = &model.WfParty{}
	var workitem = &model.WfWorkitem{}
	workitem, err = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if permanent {
		task := NewManualAssignmentTask(workitem.CaseID, workitem.TransitionID, partyID)
		HandleManualAssignmentTask(context.Background(), task)
	}
	var workitemAssignment = &model.WfWorkitemAssignment{}
	workitemAssignments, err := workitemAssignment.Find(model.NewSqlCnd().Where("workitem_id =? and party_id =?", workitemID, partyID))
	if len(workitemAssignments) > 0 {
		workitemAssignment = &workitemAssignments[0]
		workitemAssignment.Delete(workitemAssignment.ID)
	}
	//unAssignmentCallBack := workitem.Transition.UnassignmentCallback
	//callbacks.Register(workitem.Transition)
	//callbacks.SetInstance("model.WfTranstion").(model.WfTranstion).UnassignmentCallback
	return nil
}
