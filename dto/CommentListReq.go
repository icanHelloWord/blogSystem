package dto

type CommentListReq struct {
	PostID   int `json:"postId" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
	PageNum  int `json:"pageNum" binding:"required"`
}
