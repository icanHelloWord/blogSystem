package dto

import "time"

type UpdatePostsReq struct {
	ID         int       `json:"id"  binding:"required"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserId     int       `json:"userId"`
	UpdateDate time.Time `gorm:"column:update_date;autoUpdateTime" json:"update_date"` // 更新时间
}
