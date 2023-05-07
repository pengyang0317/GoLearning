package model

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null;default:'';comment:'品牌名称'" json:"name"`
	Logo string `gorm:"type:varchar(255);not null;default:'';comment:'品牌logo'" json:"logo"`
}
