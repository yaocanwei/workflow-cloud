package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeWorkitemAssignmentSet = "workitem_assignment:set"
)

func NewWorkitemAssignmentSetTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeCaseResume, payload)
}

func HandleWorkitemAssignmentSetTask(ctx context.Context, t *cmdx.Task) error {
	var hasCaseAss = false
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(workitemID)
	var workitem = &model.WfWorkitem{}
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	var caseAssignment = &model.WfCaseAssignment{}
	caseAssignments, _ := caseAssignment.Find(model.NewSqlCnd().Where("case_id =? and transition_id =?", workitem.CaseID, workitem.TransitionID))
	for _, caseAssignment := range caseAssignments {
		task := NewWorkitemAssignmentAddTask(workitemID, caseAssignment.PartyID, false)
		HandleWorkitemAssignmentAddTask(context.Background(), task)
		hasCaseAss = true
	}
	if !hasCaseAss {
		//assignmentCallback := workitem.Transition.AssignmentCallback
		//callbacks.SetInstance(assignmentCallback)
		staticAssignments := workitem.Transition.TransitionStaticAssignments
		for _, staticAssignment := range staticAssignments {
			task := NewWorkitemAssignmentAddTask(workitemID, staticAssignment.PartyID, false)
			HandleWorkitemAssignmentAddTask(context.Background(), task)
		}
	}
	return nil
}
