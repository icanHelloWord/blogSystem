package repository

import (
	"blogSystem/dto"
	"blogSystem/model"
	"blogSystem/utils"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) CommentList(commentListReq *dto.CommentListReq) []model.Comment {
	var comments []model.Comment
	query := r.db.Model(&model.Comment{})
	query.Scopes(utils.PageInfo(commentListReq.PageNum, commentListReq.PageSize))
	query.Where("post_id = ?", commentListReq.PostID).Find(&comments)
	return comments
}
func (r *commentRepository) Delete(id int) error {
	return r.db.Delete(&model.Comment{}, id).Error
}

func (r *commentRepository) FindById(id int) (model.Comment, error) {
	var comment model.Comment
	return comment, r.db.Where("id = ?", id).First(&comment).Error
}
