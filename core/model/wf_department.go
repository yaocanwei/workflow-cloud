package model

import (
	"time"
)

// WfDepartment 部门
type WfDepartment struct {
	Model
	Name      string    `json:"name"`
	Users     []WfUser  `gorm:"foreignKey:DepartmentID"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
