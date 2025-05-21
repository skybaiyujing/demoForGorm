package service

import (
	"Go_project/model"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, user *model.User) error {
	return db.Save(user).Error
}
