package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
)

// TypeWorkitemActionBegin 操作任务定义
const (
	TypeWorkitemActionBegin = "workitem_action:begin"
)

// NewWorkitemActionBeginTask 初始化
func NewWorkitemActionBeginTask(workitemID, userID, action int) *cmdx.Task {
	payload := map[string]interface{}{"workitem_id": workitemID, "user_id": userID, "action": action}
	return cmdx.NewTask(TypeWorkitemActionBegin, payload)
}

// HandleWorkitemActionBeginTask 任务开始逻辑处理
func HandleWorkitemActionBeginTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	action, err := t.Payload.GetString("action")
	fmt.Println(workitemID, userID, action)
	var workitem *model.WfWorkitem
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	var msg string
	if action == "start" {
		if workitem.State != 0 {
			msg = fmt.Sprintf("只有启用了的任务才能开始")
			panic(msg)
		}
		if !workitem.OwnBy(userID) {
			panic("你没有被指派到这个任务")
		}
	}
	if action == "finish" || action == "cancel" {
		if workitem.State == 1 && workitem.HoldingUserID != userID {
			panic("你非当前任务相关人")
		}
		if workitem.State == 0 {
			if action == "cancel" {
				msg = fmt.Sprintf("任务开始了才能取消")
				panic(msg)
			}
			if !workitem.OwnBy(userID) {
				panic("你没有被指派到这个任务")
			}
			workitem.HoldingUserID = userID
			err = workitem.Updates(map[string]interface{}{"id": workitem.ID})
			if err != nil {
			}
		} else {
			panic("任务须为启用或开始状态")
		}
	}
	return nil
}
