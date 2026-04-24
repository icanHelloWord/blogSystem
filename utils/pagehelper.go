package utils

import "gorm.io/gorm"

func PageInfo(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}
		if pageSize <= 0 {
			pageSize = 5
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}
