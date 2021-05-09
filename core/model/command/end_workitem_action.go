package command

import (
	"context"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeWorkitemActionEnd = "workitem_action:end"
)

func NewWorkitemActionEndTask(workitemID, userID int, action string) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
		"user_id":     userID,
		"action":      action,
	}
	return cmdx.NewTask(TypeWorkitemActionEnd, payload)
}

func HandleWorkitemActionEndTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	action, err := t.Payload.GetString("action")
	if err != nil {
		return err
	}
	fmt.Println(workitemID, userID, action)
	return nil
}
