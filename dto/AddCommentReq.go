package dto

import "time"

type AddCommentReq struct {
	Content    string    `json:"content" binding:"required"` // 评论内容
	UserID     int       `json:"userId" binding:"required"`
	PostID     int       `json:"postId" binding:"required"`
	CreateDate time.Time `json:"createDate,"`
	UpdateDate time.Time `json:"updateDate,"`
}
