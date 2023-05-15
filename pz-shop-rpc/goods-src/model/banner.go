package model

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(255);not null;default:'';comment:'图片'"`
	Url   string `gorm:"type:varchar(255);not null;default:'';comment:'链接'"`
	Index int32  `gorm:"type:int;not null;default:0;comment:'排序';uniqueIndex:idx_index"`
}
