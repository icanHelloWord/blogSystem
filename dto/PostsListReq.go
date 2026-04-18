package dto

type PostsListReq struct {
	Title    string `json:"title"`
	PageSize int    `json:"pageSize" binding:"required"`
	PageNum  int    `json:"pageNum" binding:"required"`
}
