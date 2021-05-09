package model

import (
	"time"
)

// WfRole 角色
type WfRole struct {
	Model
	Name      string    `json:"name"`
	Users     []WfUser  `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
