package command

import (
	"context"
	"flowpipe-server/core/model"
	"flowpipe-server/pkg/cmdx"
)

const (
	TypeAutomaticTransitionsSweep = "automatic_transitions:sweep"
)

// NewAutomaticTransitionsSweepTask NewAutomaticTransitionsSweepTask
func NewAutomaticTransitionsSweepTask(caseID int) *cmdx.Task {
	payload := map[string]interface{}{
		"case_id": caseID,
	}
	return cmdx.NewTask(TypeAutomaticTransitionsSweep, payload)
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 3:54 AM 2021/3/26
 * @Param
 * @return
 **/
func HandleAutomaticTransitionsSweepTask(ctx context.Context, t *cmdx.Task) error {
	caseID, err := t.Payload.GetInt("case_id")
	if err != nil {
		return err
	}

	var wfWorkitem = &model.WfWorkitem{}
	wfWorkitems, err := wfWorkitem.Find(model.NewSqlCnd().
		Joins("left join wf_transitions on wf_transitions.id = wf_workitems.transition_id").
		Where("wf_workitems.case_id =? and wf_workitems.state = ? and wf_transitions.trigger_type", caseID, 0, 1))
	if err != nil {
		return err
	}

	task := NewTransitionEnableTask(caseID)
	err = HandleTransitionEnableTask(context.Background(), task)
	if err != nil {
		return err
	}

	done := false
	for {
		done = true
		task = NewPFinishTask(caseID)
		result, _ := HandlePFinishTask(context.Background(), task)
		if result {
			continue
		}
		for _, workitem := range wfWorkitems {
			fireTransitonInternalTask := NewTransitionInternalFireTask(workitem.ID)
			HandleTransitionInternalFireTask(context.Background(), fireTransitonInternalTask)
			done = false
		}
		enableTransitionTask := NewTransitionEnableTask(caseID)
		HandleTransitionEnableTask(context.Background(), enableTransitionTask)
		if done {
			break
		}
	}

	return nil
}
