package model

import (
	"time"
)

// WfTag 标签
type WfTag struct {
	Model
	Name      string    `json:"name"`
	Users     []WfUser  `gorm:"foreignKey:TagID"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
