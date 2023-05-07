package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID int32 `gorm:"primary_key;type:int" json:"id"`
	// CreateAt  *time.Time     `gorm:"column:add_time" json:"-"`
	// UpdateAt  *time.Time     `gorm:"column:update_time" json:"-"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdateAt  time.Time      `gorm:"column:update_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"update_at,omitempty"`
	DeleteAt  gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}
