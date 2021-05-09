package model

import (
	db "flowpipe-server/core/db"
	"time"
)

// WfToken 动态对象，可以从一个库所移动到另一个库所
type WfToken struct {
	Model
	WorkflowID     int    `json:"workflow_id"`
	CaseID         int    `json:"case_id"`
	TargetableType string `json:"targetable_type"`
	TargetableID   int    `json:"targetable_id"`
	PlaceID        int    `json:"place_id"`
	// 0-free, 1-locked, 2-canceled, 3-consumed
	State            int       `json:"state"`
	LockedWorkitemID int       `json:"locked_workitem_id"`
	ProducedAt       time.Time `gorm:"column:produced_at" json:"produced_at"`
	LockedAt         time.Time `gorm:"column:locked_at" json:"locked_at"`
	CanceledAt       time.Time `gorm:"column:canceled_at" json:"canceled_at"`
	ConsumedAt       time.Time `gorm:"column:consumed_at" json:"consumed_at"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `json:"column:updated_at"`
}

// Create 创建token
func (wfToken *WfToken) Create() (*WfToken, error) {
	err := db.WfDb.Create(&wfToken).Error
	return wfToken, err
}

// Updates 更新token属性
func (wfToken *WfToken) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(wfToken).Error
}

/**
 * @Author canweiyao
 * @Description //TODO
 * @Date 7:12 PM 2021/4/4
 * @Param
 * @return
 **/
func (wfToken *WfToken) FindBy(cond *SqlCond) (*WfToken, error) {
	err := cond.FindBy(db.WfDb, wfToken)
	if err != nil {
		return &WfToken{}, err
	}
	return wfToken, nil
}

/**
 * @Author canweiyao
 * @Description //TODO 
 * @Date 7:11 PM 2021/4/4
 * @Param 
 * @return 
 **/
func (wfToken *WfToken) Find(cond *SqlCond) ([]WfToken, error) {
	wfTokenList := make([]WfToken, 0)
	err := cond.Find(db.WfDb, &wfTokenList)
	return wfTokenList, err
}
