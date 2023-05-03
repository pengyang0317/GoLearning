package utils

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `json:"id" db:"user_id"`
	Name     string `json:"name" db:"user_name"`
	Password string `json:"-" db:"-"`
}

type UserProfile struct {
	UserID int    `json:"user_id" db:"user_id"`
	Email  string `json:"email" db:"user_email"`
}

func MergeTags(dst, src interface{}) {
	dstType := reflect.TypeOf(dst).Elem()
	srcType := reflect.TypeOf(src).Elem()

	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		srcField, ok := srcType.FieldByName(dstField.Name)
		if ok {
			dstTag := dstField.Tag.Get("db")
			srcTag := srcField.Tag.Get("db")
			if dstTag == "" {
				dstField.Tag = srcField.Tag
			} else if srcTag != "" {
				dstField.Tag = reflect.StructTag(fmt.Sprintf("%s;%s", dstTag, srcTag))
			}
		}
	}
}
