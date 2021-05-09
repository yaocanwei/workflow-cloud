package command

import (
	"context"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

const (
	TypeWorkitemActionHandle = "workitem_action:handle"
)

/**
 * @Author canweiyao
 * @Description //TODO 
 * @Date 6:15 PM 2021/4/4
 * @Param 
 * @return 
 **/
func NewWorkitemActionHandleTask(workitemID, userID int, action string) *cmdx.Task {
	payload := map[string]interface{}{
		"workitem_id": workitemID,
		"user_id":     userID,
		"action":      action,
	}
	return cmdx.NewTask(TypeWorkitemActionHandle, payload)
}

/**
 * @Author canweiyao
 * @Description //TODO 
 * @Date 6:13 PM 2021/4/4
 * @Param 
 * @return 
 **/
func HandleWorkitemActionHandleTask(ctx context.Context, t *cmdx.Task) error {
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
	fmt.Println(userID, workitemID, action)
	return nil
}
