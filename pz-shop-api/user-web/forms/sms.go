package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`
	Code   string `form:"code" json:"code" binding:"required,min=5,max=5"`
}
