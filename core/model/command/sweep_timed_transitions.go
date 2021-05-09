package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
	"time"
)

const (
	TypeTimedTransitionsSweep = "timed_transitions:sweep"
)

// NewTimedTransitionsSweepTask NewTimedTransitionsSweepTask
func NewTimedTransitionsSweepTask() *cmdx.Task {
	return cmdx.NewTask(TypeAutomaticTransitionsSweep, nil)
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 3:38 AM 2021/3/26
 * @Param
 * @return
 **/
func HandleTimedTransitionsSweepTask(ctx context.Context, t *cmdx.Task) error {
	var wfWorkitem = &model.WfWorkitem{}
	wfWorkitemList, err := wfWorkitem.Find(model.NewSqlCnd().Where("trigger_time <= ?", time.Now()))
	if err != nil {
		return err
	}
	for _, wfWorkitem := range wfWorkitemList {
		task1 := NewTransitionInternalFireTask(wfWorkitem.ID)
		HandleTransitionInternalFireTask(context.Background(), task1)

		task2 := NewAutomaticTransitionsSweepTask(wfWorkitem.CaseID)
		HandleAutomaticTransitionsSweepTask(context.Background(), task2)
	}
	return nil
}
