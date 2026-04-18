package repository

import (
	"blogSystem/dto"
	"blogSystem/model"

	"gorm.io/gorm"
)

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return &postsRepository{
		db: db,
	}
}

func (postsRepository *postsRepository) Create(posts *model.Posts) error {
	result := postsRepository.db.Create(posts)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postsRepository *postsRepository) PostsList(postsListReq *dto.PostsListReq) []model.Posts {
	var posts []model.Posts
	query := postsRepository.db.Model(&model.Posts{})
	query.Scopes(pageInfo(postsListReq.PageNum, postsListReq.PageSize))
	query.Where("is_deleted = ?", 1)
	if postsListReq.Title != "" {
		// query.Where("title like ?", gorm.Expr("concat('%', ?, '%')", postsListReq.Title))
		query.Where("title like concat('%',?,'%')", postsListReq.Title)
	}
	query.Find(&posts)
	// postsRepository.db.Scopes(pageInfo(postsListReq.PageNum, postsListReq.PageSize)).
	// 	Where("is_deleted = ?", 1).
	// 	Find(&posts)
	return posts

}

func (postsRepository *postsRepository) Delete(id int) error {
	// result := db.Model(&models.User{}).Where("ID = ?", 11).Updates(&user)
	result := postsRepository.db.Model(&model.Posts{}).
		Where("id = ?", id).Update("is_deleted", 0)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (postsRepository *postsRepository) FindById(id int) (model.Posts, error) {
	var posts model.Posts
	result := postsRepository.db.Model(&model.Posts{}).Where("id = ?", id).Find(&posts)
	if result.Error != nil {
		return posts, result.Error
	}
	return posts, nil
}

func (postsRepository *postsRepository) Update(postsReq *dto.UpdatePostsReq) error {

	result := postsRepository.db.Model(&model.Posts{}).
		Where("id = ?", postsReq.ID).
		Updates(postsReq)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func pageInfo(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}
		if pageSize <= 0 {
			pageSize = 5
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}

}
