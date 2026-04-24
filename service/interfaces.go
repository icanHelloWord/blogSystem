package service

import (
	"blogSystem/dto"
	"blogSystem/model"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Login(c *gin.Context, req *dto.UserLoginReq) (string, error)

	Register(c *gin.Context, req *dto.UserLoginReq) error
}

type PostsService interface {
	AddPosts(posts *dto.AddPosts) error

	PostsList(postsListReq *dto.PostsListReq) []model.Posts

	DeletePosts(postsId int, userId int) error

	UpdatePosts(posts *dto.UpdatePostsReq) error
}

type CommentService interface {
	AddComment(comment *dto.AddCommentReq) error

	CommentList(commentListReq *dto.CommentListReq) []model.Comment

	Delete(id int, userID int) error
}
