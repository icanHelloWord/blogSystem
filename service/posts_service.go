package service

import (
	"blogSystem/dto"
	"blogSystem/model"
	"blogSystem/repository"
	"errors"
)

type postsService struct {
	postsRepository repository.PostsRepository
}

func NewPostsService(postsRepository repository.PostsRepository) PostsService {
	return &postsService{
		postsRepository: postsRepository,
	}
}

func (postsService *postsService) AddPosts(postsDto *dto.AddPosts) error {
	var posts model.Posts
	posts.Title = postsDto.Title
	posts.Content = postsDto.Content
	posts.UserID = postsDto.UserID
	if err := postsService.postsRepository.Create(&posts); err != nil {
		return err
	}

	return nil
}

func (postsService *postsService) PostsList(postsListReq *dto.PostsListReq) []model.Posts {
	return postsService.postsRepository.PostsList(postsListReq)
}

func (postsService *postsService) DeletePosts(postsId int, userId int) error {
	posts, err := postsService.postsRepository.FindById(postsId)
	if err != nil {
		return err
	}
	if posts.UserID != userId {
		return errors.New("只能删除自己的博客")
	}
	return postsService.postsRepository.Delete(postsId)
}

func (postsService *postsService) UpdatePosts(postsReq *dto.UpdatePostsReq) error {

	posts, err := postsService.postsRepository.FindById(postsReq.ID)
	if err != nil {
		return err
	}
	if posts.UserID != postsReq.UserId {
		return errors.New("只能更新自己的博客")
	}
	return postsService.postsRepository.Update(postsReq)
}
