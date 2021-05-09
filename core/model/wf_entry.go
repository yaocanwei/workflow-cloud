package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfEntry 流程实例
type WfEntry struct {
	Model
	UserID      int            `json:"user_id"`
	WorkitemID  int            `json:"workitem_id"`
	Payload     string         `json:"payload" gorm:"type:text"`
	FormID      int            `json:"form_id"`
	FieldValues []WfFieldValue `json:"field_values"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"updated_at" json:"updated_at"`
	// TODO index
}

// Create 创建entry
func (wfEntry *WfEntry) Create() (*WfEntry, error) {
	err := db.WfDb.Create(&wfEntry).Error
	return wfEntry, err
}

// FindOrCreate 查找或创建
func (wfEntry *WfEntry) FindOrCreate(cond *SqlCond) (*WfEntry, error) {
	foundedRecord, err := wfEntry.FindBy(cond)
	if err == nil {
		return foundedRecord, nil
	}
	wfEntry, err = wfEntry.Create()
	return wfEntry, err
}

// Updates 更新entry
func (wfEntry *WfEntry) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(wfEntry).Error
}

// FindBy 根据条件查找记录
func (wfEntry *WfEntry) FindBy(cond *SqlCond) (*WfEntry, error) {
	err := cond.FindBy(db.WfDb, wfEntry)
	if err != nil {
		return &WfEntry{}, err
	}
	return wfEntry, nil
}
