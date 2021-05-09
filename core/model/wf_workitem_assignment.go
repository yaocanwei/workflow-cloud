package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfWorkitemAssignment 指派任务
type WfWorkitemAssignment struct {
	Model
	PartyID    int       `json:"party_id"`
	WorkitemID int       `json:"workitem_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	// TODO index
}

// FindByPartyID  根据组织ID查找workitem
func (workitemAssignment *WfWorkitemAssignment) FindByPartyID(workitemID, partyID int) (*WfWorkitemAssignment, error) {
	var err error
	if err = db.WfDb.Where("workitem_id = ? and party_id = ?", workitemID, partyID).Find(&workitemAssignment).Error; err != nil {
		return nil, err
	}
	return workitemAssignment, nil
}

/**
 * @Author canweiyao
 * @Description  获取指派任务
 * @Date 8:34 PM 2021/4/4
 * @Param
 * @return
 **/
func (workitemAssignment *WfWorkitemAssignment) Find(cond *SqlCond) ([]WfWorkitemAssignment, error) {
	workitemAssignmentList := make([]WfWorkitemAssignment, 0)
	err := cond.Find(db.WfDb, &workitemAssignmentList)
	return workitemAssignmentList, err
}

/**
 * @Author canweiyao
 * @Description  返回单条workitemAssignment
 * @Date 8:35 PM 2021/4/4
 * @Param
 * @return
 **/
func (workitemAssignment *WfWorkitemAssignment) FindBy(cond *SqlCond) (*WfWorkitemAssignment, error) {
	err := cond.FindBy(db.WfDb, workitemAssignment)
	if err != nil {
		return &WfWorkitemAssignment{}, err
	}
	return workitemAssignment, nil
}

// Create 创建 workitemassignment
func (workitemAssignment *WfWorkitemAssignment) Create() (*WfWorkitemAssignment, error) {
	err := db.WfDb.Create(&workitemAssignment).Error
	return workitemAssignment, err
}

// Delete 删除 workitemAssignment
func (workitemAssignment *WfWorkitemAssignment) Delete(id int) error {
	return db.WfDb.Where("id = ?", id).Delete(&workitemAssignment).Error
}
