package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
)

const (
	TypeWorkitemStart = "workitem:start"
)

// NewWorkitemStartTask NewWorkitemStartTask
func NewWorkitemStartTask(workitemID, userID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
		"user_id":     userID,
	}
	return cmdx.NewTask(TypeWorkitemStart, payload)
}

// HandleWorkitemStartTask HandleWorkitemStartTask
func HandleWorkitemStartTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	var workitem = &model.WfWorkitem{}
	workitem, err = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if err != nil {
		return err
	}
	workitem.Updates(map[string]interface{}{"state": 1, "holding_user_id": userID})
	var arc = model.WfArc{}
	arcs, _ := arc.Find(model.NewSqlCnd().Where("transition_id =? and direction =?", workitem.TransitionID, 0))
	for _, arc := range arcs {
		task := NewTokenLockTask(workitem.CaseID, arc.PlaceID, workitemID)
		HandleTokenLockTask(context.Background(), task)
	}
	return nil
}
