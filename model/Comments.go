package model

// id int(11) unsigned NOT NULL
// content varchar(255) NULL评论内容
// user_id int(11) NULL
// post_id nt(11) NULL
// create_date timestamp NULL
// update_date timestamp NULL
// is_deleted tinyint(1) NULL0:删除,1:未删除
import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement" json:"id"`
	Content    string    `gorm:"column:content;type:varchar(255);default:NULL;index" json:"content,omitempty"` // 评论内容
	UserID     int       `gorm:"column:user_id;type:int(11);default:NULL;index" json:"userId,omitempty"`
	PostID     int       `gorm:"column:post_id;type:int(11);default:NULL;index" json:"postId,omitempty"`
	CreateDate time.Time `gorm:"column:create_date;type:timestamp;default:NULL" json:"createDate,omitempty"`
	UpdateDate time.Time `gorm:"column:update_date;type:timestamp;default:NULL" json:"updateDate,omitempty"`
	IsDeleted  int       `gorm:"column:is_deleted;type:tinyint(1);default:0;index" json:"isDeleted,omitempty"` // 0:删除,1:未删除
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
}

// 软删除钩子
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	if tx.Statement.Unscoped {
		return
	}

	now := time.Now()
	tx.Statement.SetColumn("is_deleted", 1)
	tx.Statement.SetColumn("update_date", now)
	return
}

// 创建评论
func CreateComment(db *gorm.DB, content string, userID, postID int) error {
	now := time.Now()
	comment := &Comment{
		Content:    content,
		UserID:     userID,
		PostID:     postID,
		CreateDate: now,
		UpdateDate: now,
		IsDeleted:  1,
	}
	return db.Create(comment).Error
}

// 查询未删除的评论
func GetCommentsByPostID(db *gorm.DB, postID int) ([]Comment, error) {
	var comments []Comment
	err := db.Where("post_id = ? AND is_deleted = 0", postID).Find(&comments).Error
	return comments, err
}

// 软删除
func SoftDeleteComment(db *gorm.DB, id uint) error {
	return db.Delete(&Comment{}, id).Error
}

// 指针转换辅助函数
func boolPtr(b bool) *bool {
	return &b
}
