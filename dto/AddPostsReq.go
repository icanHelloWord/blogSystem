package dto

type AddPosts struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  int    `json:"userId"`
}
