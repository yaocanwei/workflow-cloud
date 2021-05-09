package model

import(
    "time"
)

type WfDemoTarget struct {
    Model
    Name string `json:"name"`
    Description string `json:"description"`
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
    UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
