package repository

import (
	"blogSystem/dto"
	"blogSystem/model"
)

type UserRepository interface {
	Create(user *model.User) error

	GetUserByUserName(username string) (*dto.UserDto, error)
}

type PostsRepository interface {
	Create(posts *model.Posts) error

	PostsList(postsListReq *dto.PostsListReq) []model.Posts

	Delete(id int) error

	FindById(id int) (model.Posts, error)

	Update(posts *dto.UpdatePostsReq) error
}
