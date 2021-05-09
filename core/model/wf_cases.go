package model

import (
	db "flowpipe-server/core/db"
	"strconv"
	"time"
)

// WfCase 流程实例
type WfCase struct {
	Model
	WorkflowID int `json:"workflow_id"`
	Workflow   WfWorkflow  `json:"workflow" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:WorkflowID"`
	// point to target type of Application.
	TargetableType string `json:"targetable_type"`
	TargetableID   int    `json:"targetable_id"`
	// 0-created, 1-active, 2-suspended, 3-canceled, 4-finished
	State           int                `json:"state"`
	CaseAssignments []WfCaseAssignment `gorm:"foreignKey:CaseID"`
	CreatedAt       time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time          `gorm:"updated_at" json:"updated_at"`
}

// ChangeState 更新流程实例状态
func (wfCase *WfCase) ChangeState(param string) error {
	state := 0
	switch param {
	case "created":
		state = 0
	case "active":
		state = 1
	case "suspended":
		state = 2
	case "canceled":
		state = 3
	default:
		state = 4
	}

	if err := db.WfDb.Model(&wfCase).Update("state", state).Error; err != nil {
		return err
	}
	return nil
}

// Create 创建
func (wfCase *WfCase) Create() (*WfCase, error) {
	err := db.WfDb.Create(&wfCase).Error
	return wfCase, err
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:05 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfCase *WfCase) Find(cond *SqlCond) ([]WfCase, error) {
	wfCaseList := make([]WfCase, 0)
	err := cond.Find(db.WfDb, &wfCaseList)
	return wfCaseList, err
}

/**
 * @Author canweiyao
 * @Description 返回case
 * @Date 7:04 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfCase *WfCase) FindBy(cond *SqlCond) (*WfCase, error) {
	//err := cond.FindBy(db.WfDb.Preload("CaseAssignments").Preload("Workflow.Transitions"), wfCase)
	err := cond.FindBy(db.WfDb, wfCase)
	if err != nil {
		return &WfCase{}, err
	}
	return wfCase, nil
}

// Delete 根据id删除流程实例
func (wfCase *WfCase) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&wfCase).Error
}

// Updates 更新任务
func (wfCase *WfCase) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(&wfCase).Error
}

// Count 统计case
func (wfCase *WfCase) Count() (count int, err error) {
	if err = db.WfDb.Model(&wfCase).Unscoped().Count(&count).Error; err != nil {
		return
	}
	return
}

// CanFire 判断能否fire
func (wfCase *WfCase) CanFire(transitionID int) bool {
	var arc = &WfArc{}
	arcs, _ := arc.Find(NewSqlCnd().Where("transition_id =? and direction =?", transitionID, 0))
	if len(arcs) == 0 {
		return false
	}
	for _, arc := range arcs {
		var token = &WfToken{}
		tokens, _ := token.Find(NewSqlCnd().Where("place_id =? and case_id =? and state =?", arc.PlaceID, wfCase.ID, 0))
		if len(tokens) > 0 {
			return true
		}
	}
	return false
}

// Name case name
func (wfCase *WfCase) Name() string {
	return "Case->" + strconv.Itoa(wfCase.ID)
}
