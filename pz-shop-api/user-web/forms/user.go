package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord  string `form:"password" json:"password" binding:"required,min=6,max=16"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type UpdateUserForm struct {
	Name     string `form:"name" json:"name" binding:"required,min=3,max=10"`
	Gender   string `form:"gender" json:"gender" binding:"required,oneof=female male"`
	Birthday string `form:"birthday" json:"birthday" binding:"required,datetime=2023-05-01"`
}
