package command

import (
	"context"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeTransitionInternalFire = "transition_internal:fire"
)

func NewTransitionInternalFireTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
	}
	return cmdx.NewTask(TypeTransitionInternalFire, payload)
}

func HandleTransitionInternalFireTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Println(workitemID)
	return nil
}
