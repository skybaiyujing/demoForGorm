package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	// 使用 gorm 标签来指定字段的属性
	ID       int `gorm:"primaryKey;autoIncrement"` // 主键
	Name     string
	Age      int
	Birthday time.Time
	Email    string //`gorm:"unique"` // 唯一键
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Name == "admin" {
		return errors.New("admin user not allowed to update")
	}
	return nil
}
