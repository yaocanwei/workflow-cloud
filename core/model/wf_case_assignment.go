package model

import (
	db "flowpipe-server/core/db"
	"time"
)

// WfCaseAssignment
type WfCaseAssignment struct {
	Model
	// 流程定义ID
	CaseID       int       `json:"case_id"`
	TransitionID int       `json:"transition_id"`
	PartyID      int       `json:"party_id"`
	CreatedAt    time.Time `gorm:"column:create_at" json:"create_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	// TODO: 设置index
}

/**
 * @Author canweiyao
 * @Description  获取指派任务
 * @Date 7:01 PM 2021/4/4
 * @Param
 * @return
 **/
func (caseAssignment *WfCaseAssignment) Find(cond *SqlCond) ([]WfCaseAssignment, error) {
	wfCaseAssList := make([]WfCaseAssignment, 0)
	err := cond.Find(db.WfDb, &wfCaseAssList)
	return wfCaseAssList, err
}


/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:02 PM 2021/4/4
 * @Param
 * @return
 **/
func (caseAssignment *WfCaseAssignment) FindBy(cond *SqlCond) (*WfCaseAssignment, error) {
	err := cond.FindBy(db.WfDb, caseAssignment)
	if err != nil {
		return &WfCaseAssignment{}, err
	}
	return caseAssignment, nil
}


// Create 手动指派人
func (caseAssignment *WfCaseAssignment) Create() (*WfCaseAssignment, error) {
	err := db.WfDb.Create(&caseAssignment).Error
	return caseAssignment, err
}

// Delete 根据id删除
func (caseAssignment *WfCaseAssignment) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&caseAssignment).Error
}
