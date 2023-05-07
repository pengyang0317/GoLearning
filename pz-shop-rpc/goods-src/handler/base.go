package handler

import "gorm.io/gorm"

func Paginate(PageNum, PageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if PageNum == 0 {
			PageNum = 1
		}
		switch {
		case PageSize > 100:
			PageSize = 100
		case PageSize <= 0:
			PageSize = 10
		}

		offset := (PageNum - 1) * PageSize
		return db.Offset(offset).Limit(PageSize)
	}
}
