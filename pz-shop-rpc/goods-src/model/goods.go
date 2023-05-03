package model

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(50);not null;default:'';comment:'分类名称'" json:"name"`
	ParentCategoryID int32       `gorm:"comment:'父分类ID'" json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignkey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1;comment:'分类级别'" json:"level"`
	IsTab            bool        `gorm:"type:bool;not null;default:false;comment:'是否导航'" json:"is_tab"`
}
