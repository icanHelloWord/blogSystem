// idint(11) unsigned NOT NULL
// titlevarchar(50) NULL标题
// contentvarchar(255) NULL内容
// user_idint(11) NULL用户id
// create_datetimestamp NULL
// update_datetimestamp NULL
// id_deletedtinyint(1) NULL0:删除,1:未删除

package model

import (
	"time"

	"gorm.io/gorm"
)

// Posts 文章表实体
type Posts struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement;unsigned" json:"id"`
	Title      string    `gorm:"column:title;type:varchar(50);not null" json:"title"`
	Content    string    `gorm:"column:content;type:varchar(255)" json:"content"`
	UserID     int       `gorm:"column:user_id;not null;index" json:"userId"`
	CreateDate time.Time `gorm:"column:create_date;autoCreateTime" json:"createDate"`
	UpdateDate time.Time `gorm:"column:update_date;autoUpdateTime" json:"updateDate"`
	IsDeleted  int8      `gorm:"column:is_deleted;type:tinyint(1);default:1" json:"isDeleted"` // 删除标记 0:删除,1:未删除
}

// TableName 指定表名
func (Posts) TableName() string {
	return "Posts"
}

// 删除状态常量
const (
	ArticleDeleted    int8 = 0
	ArticleNotDeleted int8 = 1
)

// 以下是可选的钩子方法
func (a *Posts) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreateDate = time.Now()
	a.UpdateDate = time.Now()
	a.IsDeleted = ArticleNotDeleted
	return
}

func (a *Posts) BeforeUpdate(tx *gorm.DB) (err error) {
	a.UpdateDate = time.Now()
	return
}

// SoftDelete 软删除文章
func (a *Posts) SoftDelete() {
	a.IsDeleted = ArticleDeleted
	a.UpdateDate = time.Now()
}

// IsActive 检查文章是否活跃（未删除）
func (a *Posts) IsActive() bool {
	if a.IsDeleted == ArticleNotDeleted {
		return true
	}
	return false
}
