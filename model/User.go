// idint(11) NOT NULL
// usernamevarchar(25) NULL用户名
// passwordvarchar(255) NULL密码
// emailvarchar(25) NULL邮箱
// create_datetimestamp NULL
// update_datetimestamp NULL
// is_deletedtinyint(1) NULL0:删除,1:未删除
package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	Deleted    int8 = 0 // 已删除
	NotDeleted int8 = 1 // 未删除
)

type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                  // 用户ID
	Username   string    `gorm:"column:username;type:varchar(25)" json:"username"`              // 用户名
	Password   string    `gorm:"column:password;type:varchar(255)" json:"-"`                    // 密码（JSON不返回）
	Email      string    `gorm:"column:email;type:varchar(25)" json:"email"`                    // 邮箱
	CreateDate time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date"`          // 创建时间
	UpdateDate time.Time `gorm:"column:update_date;autoUpdateTime" json:"update_date"`          // 更新时间
	IsDeleted  int8      `gorm:"column:is_deleted;type:tinyint(1);default:1" json:"is_deleted"` // 删除标记 0:删除,1:未删除
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// 以下是可选的方法

// BeforeCreate 创建前的钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateDate = time.Now()
	u.UpdateDate = time.Now()
	u.IsDeleted = NotDeleted
	return
}

// BeforeUpdate 更新前的钩子
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdateDate = time.Now()
	return
}

// IsActive 检查用户是否活跃（未删除）
func (u *User) IsActive() bool {
	return u.IsDeleted == NotDeleted
}

// SoftDelete 软删除用户
func (u *User) SoftDelete() {
	u.IsDeleted = Deleted
	u.UpdateDate = time.Now()
}
