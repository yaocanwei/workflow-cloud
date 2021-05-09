package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfFieldValue 字段值
type WfFieldValue struct {
	Model
	FormID    int       `json:"form_id"`
	FieldID   int       `json:"field_id"`
	Value     string    `json:"value" gorm:"type:text"`
	EntryID   int       `json:"entry_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
	// TODO index
}

// Create 创建form field
func (wfFieldValue *WfFieldValue) Create() (*WfFieldValue, error) {
	err := db.WfDb.Create(&wfFieldValue).Error
	return wfFieldValue, err
}

// Updates 更新form field
func (wfFieldValue *WfFieldValue) Updates(condition map[string]interface{}) error {
	return db.WfDb.Where(condition).Updates(wfFieldValue).Error
}

// FindBy 根据给定条件查找记录
func (wfFieldValue *WfFieldValue) FindBy(cond *SqlCond) (*WfFieldValue, error) {
	err := cond.FindBy(db.WfDb, wfFieldValue)
	if err != nil {
		return &WfFieldValue{}, err
	}
	return wfFieldValue, nil
}
