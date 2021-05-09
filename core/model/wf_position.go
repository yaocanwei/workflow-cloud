package model

import (
	"time"
)

// WfPosition 岗位
type WfPosition struct {
	Model
	Name      string    `json:"name"`
	Users     []WfUser  `gorm:"foreignKey:PositionID"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
