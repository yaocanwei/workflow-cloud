package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"fmt"
	"time"
)

// TypeWorkitemCancel 取消任务变量定义
const (
	TypeWorkitemCancel = "workitem:cancel"
)

// NewWorkitemCancelTask 初始化取消任务
func NewWorkitemCancelTask(workitemID int) *cmdx.Task {
	payload := map[string]interface{}{"workitem_id": workitemID}
	return cmdx.NewTask(TypeWorkitemCancel, payload)
}

// HandleWorkitemCancelTask 任务取消逻辑处理
func HandleWorkitemCancelTask(ctx context.Context, t *cmdx.Task) error {
	workitemID, err := t.Payload.GetInt("workitem_id")
	if err != nil {
		return err
	}
	fmt.Print(workitemID)
	var workitem *model.WfWorkitem
	workitem, _ = workitem.FindBy(model.NewSqlCnd().Where("id =?", workitemID))
	if workitem.State != 1 {
		panic("非开始状态任务,无法取消")
	}
	// TODO: add transaction
	workitem.State = 2
	workitem.CanceledAt = time.Now()
	workitem.Updates(map[string]interface{}{"id": workitem.ID})
	t = NewTokenReleaseTask(workitem.ID)
	err = HandleTokenReleaseTask(ctx, t)
	t = NewAutomaticTransitionsSweepTask(workitem.CaseID)
	err = HandleAutomaticTransitionsSweepTask(ctx, t)
	return nil
}
