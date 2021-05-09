package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeMessageTransitionFire = "message_transition:fire"
)

func NewMessageTransitionFireTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeMessageTransitionFire, payload)
}

func HandleMessageTransitionFireTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(workitemID)
	var workitem = &model.WfWorkitem{}
	workitem, err = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if err != nil {
		return err
	}
	if workitem.Transition.TriggerType != 2 {
		panic("Transition #{workitem.transition.name} is not message triggered")
	}
	//TODO: add transaction
	task := NewTransitionInternalFireTask(workitemID)
	HandleTransitionInternalFireTask(context.Background(), task)
	task = NewAutomaticTransitionsSweepTask(workitem.CaseID)
	HandleAutomaticTransitionsSweepTask(context.Background(), task)
	return nil
}
