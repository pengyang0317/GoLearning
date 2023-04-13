package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 自定义验证器
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^1[23456789]\d{9}$`, mobile)
	return ok
}
