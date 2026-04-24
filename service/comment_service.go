package service

import (
	"blogSystem/dto"
	"blogSystem/model"
	"blogSystem/repository"
	"errors"
	"time"
)

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &commentService{commentRepository: commentRepository}
}

func (s *commentService) AddComment(comment *dto.AddCommentReq) error {
	com := model.Comment{
		Content:    comment.Content,
		UserID:     comment.UserID,
		PostID:     comment.PostID,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		IsDeleted:  1,
	}
	return s.commentRepository.Create(&com)
}

func (s *commentService) CommentList(commentListReq *dto.CommentListReq) []model.Comment {
	return s.commentRepository.CommentList(commentListReq)
}

func (s *commentService) Delete(id int, userID int) error {
	comment, err := s.commentRepository.FindById(id)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		err = errors.New("您没有权限删除该评论")
	}
	return s.commentRepository.Delete(id)
}
