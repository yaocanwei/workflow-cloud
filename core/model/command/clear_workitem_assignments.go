package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeWorkitemAssignmentClear = "workitem_assignment:clear"
)

func NewWorkItemAssignmentClearTask(workitemID int, permanent bool) *cmdx.Task {
	payload := map[string]interface{}{"workitem_id": workitemID, "permanent": permanent}
	return cmdx.NewTask(TypeWorkitemAssignmentClear, payload)
}

func HandleWorkItemAssignmentClearTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	permanent, err := t.Payload.GetBool("permanent")
	fmt.Print(workitemID, permanent)
	var workitem *model.WfWorkitem
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if permanent {
		t = NewManualAssignmentClearTask(workitem.CaseID, workitem.TransitionID)
		HandleManualAssignmentClearTask(ctx, t)
	}
	for _, workitemAssignment := range workitem.WorkitemAssignments {
		workitemAssignment.Delete(workitemAssignment.ID)
	}
	return nil
}
