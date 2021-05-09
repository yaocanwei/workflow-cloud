package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeManualAssignmentClear = "manual_assignment:clear"
)

func NewManualAssignmentClearTask(caseID, transitionID int) *cmdx.Task {
	payload := map[string]interface{}{"case_id": caseID, "transition_id": transitionID}
	return cmdx.NewTask(TypeManualAssignmentClear, payload)
}

func HandleManualAssignmentClearTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	transitionID, err := t.Payload.GetInt("transition_id")
	fmt.Print(caseID, transitionID)
	var wfCase *model.WfCase
	caseAssignments := wfCase.CaseAssignments
	for _, caseAssignment := range caseAssignments {
		if caseAssignment.TransitionID == transitionID {
			caseAssignment.Delete(caseAssignment.ID)
		}
	}
	return nil
}
