package model

import (
	"flowpipe-server/core/db"
	"time"
)

// WfParty 流程组
type WfParty struct {
	Model
	PartableType string       `json:"partable_type"` // departments、 groups、 roles、 positions、tag etc
	Workitems    []WfWorkitem `json:"workitems" gorm:"many2many:wf_workitem_assignments;"`
	Department   WfDepartment `json:"department" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:PartableID"` // 部门
	Position     WfPosition   `json:"position" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:PartableID"`   // 岗位
	Role         WfRole       `json:"role" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:PartableID"`       // 角色
	Tag          WfTag        `json:"tag" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:PartableID"`        // 标签
	Users        []WfUser     `gorm:"foreignKey:PartableID"`
	CreatedAt    time.Time    `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"updated_at" json:"updated_at"`
	// TODO index
}

func (party *WfParty) FindByID(id int) (*WfParty, error) {
	if err := db.WfDb.First(&party, "id =?", id).Error; err != nil {
		return nil, err
	}
	return party, nil
}
