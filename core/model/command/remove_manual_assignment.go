package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeManualAssignmentRemove = "manual_assignment:remove"
)

func NewManualAssignmentTask(caseID, transitionID, partyID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id":       caseID,
		"transition_id": transitionID,
		"party_id":      partyID,
	}
	return cmdx.NewTask(TypeManualAssignmentRemove, payload)
}

func HandleManualAssignmentTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}
	transitionID, err := t.Payload.GetInt("transition_id")
	if err != nil {
		return err
	}
	partyID, err := t.Payload.GetInt("party_id")
	if err != nil {
		return err
	}
	fmt.Println(caseID, transitionID, partyID)
	var wfCase = &model.WfCase{}
	wfCase, err = wfCase.FindBy(model.NewSqlCnd().Where("id =?", caseID))
	if err != nil {
		return err
	}
	var caseAssignment = &model.WfCaseAssignment{}
	caseAssignments, err := caseAssignment.Find(model.NewSqlCnd().Where("case_id =? and transition_id =? and party_id", caseID, transitionID, partyID))
	if err != nil {
		return err
	}
	for _, delcaseAssignment := range caseAssignments {
		//TODO: add error chain
		delcaseAssignment.Delete(delcaseAssignment.ID)
	}
	return nil
}
