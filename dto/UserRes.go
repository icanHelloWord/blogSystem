package dto

import "time"

type UserDto struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                  // 用户ID
	Username   string    `gorm:"column:username;type:varchar(25)" json:"username"`              // 用户名
	Password   string    `gorm:"column:password;type:varchar(255)" json:"-"`                    // 密码（JSON不返回）
	Email      string    `gorm:"column:email;type:varchar(25)" json:"email"`                    // 邮箱
	CreateDate time.Time `gorm:"column:create_date;autoCreateTime" json:"create_date"`          // 创建时间
	UpdateDate time.Time `gorm:"column:update_date;autoUpdateTime" json:"update_date"`          // 更新时间
	IsDeleted  int8      `gorm:"column:is_deleted;type:tinyint(1);default:1" json:"is_deleted"` // 删除标记 0:删除,1:未删除
}
