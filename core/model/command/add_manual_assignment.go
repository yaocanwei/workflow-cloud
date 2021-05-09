package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
)

// TypeManualAssignmentAdd 手动指派人
const (
	TypeManualAssignmentAdd = "manual_assignment:add"
)

// NewManualAssignmentAddTask 初始化手动指派人任务
func NewManualAssignmentAddTask(caseID, transitionID, partyID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id":       caseID,
		"transition_id": transitionID,
		"party_id":      partyID,
	}
	return cmdx.NewTask(TypeManualAssignmentAdd, payload)
}

// HandleManualAssignmentAddTask 手动指派人逻辑处理
func HandleManualAssignmentAddTask(ctx context.Context, t *cmdx.Task) error {
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

	caseAssignment := &model.WfCaseAssignment{
		CaseID:       caseID,
		TransitionID: transitionID,
		PartyID:      partyID,
	}
	caseAssignment.Create()
	return nil
}
