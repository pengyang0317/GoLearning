package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID       int32     `gorm:"primary_key;type:int" json:"id"`
	CreateAt time.Time `gorm:"column:add_time" json:"-"`
	UpdateAt time.Time `gorm:"column:update_time" json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`
	IsDeleted bool `json:"-"`
}
