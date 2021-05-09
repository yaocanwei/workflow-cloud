package model

import (
	db "flowpipe-server/core/db"
	"time"
	// "github.com/jinzhu/gorm"
)

// WfWorkflow 流程表
type WfWorkflow struct {
	Model
	Name        string        `json:"name"`
	Description string        `json:"description" gorm:"type:text"`
	IsValid     bool          `json:"is_valid"`
	CreatedBy   string        `json:"created_by"`
	UpdatedBy   string        `json:"updated_by"`
	ErrorMsg    string        `json:"error_msg" gorm:"type:text"`
	CreatedAt   time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at" json:"updated_at"`
	Places      []WfPlace     `gorm:"foreignkey:WorkflowID;association_foreignkey:ID"`
	Transitions []WfTranstion `gorm:"foreignkey:WorkflowID;association_foreignkey:ID"`
	Arcs        []WfArc       `gorm:"foreignkey:WorkflowID;association_foreignkey:ID"`
}

// Create 创建流程
func (workflow *WfWorkflow) Create() (*WfWorkflow, error) {
	err := db.WfDb.Create(&workflow).Error
	return workflow, err
}

// Find 根据id查找流程
func (workflow *WfWorkflow) Find(cond *SqlCond) ([]WfWorkflow, error) {
	wfWorkflowList := make([]WfWorkflow, 0)
	err := cond.Find(db.WfDb, &wfWorkflowList)
	return wfWorkflowList, err
}

// FindBy 根据条件查找
func (workflow *WfWorkflow) FindBy(cond *SqlCond) (*WfWorkflow, error) {
	err := cond.FindBy(db.WfDb, workflow)
	if err != nil {
		return &WfWorkflow{}, err
	}
	return workflow, nil
}

// Delete 根据id删除流程
func (workflow *WfWorkflow) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&workflow).Error
}

// Count 统计workflow
func (workflow *WfWorkflow) Count(cond *SqlCond) (count int, err error) {
	return cond.Count(db.WfDb, &workflow)
}


/**
 * @Author canweiyao
 * @Description 更新workflow
 * @Date 6:34 PM 2021/4/11
 * @Param 
 * @return 
 **/
func (wfWorkflow *WfWorkflow) Update() (err error) {
	err = db.WfDb.Save(wfWorkflow).Error
	return
}