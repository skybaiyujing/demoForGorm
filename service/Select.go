package service

import (
	"Go_project/model"
	"fmt"
	"gorm.io/gorm"
	//"time"
)

func SelectWithId(db *gorm.DB, ID string) ([]model.User, error) {
	//根据主键检索
	var user []model.User
	result := db.Where("id = ?", ID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func SelectAll(db *gorm.DB) ([]model.User, error) {
	var user []model.User
	result := db.Find(&user) // SELECT * FROM users WHERE id =
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func SelectWithField(db *gorm.DB, field string, value interface{}) ([]model.User, error) {
	var users []model.User
	result := db.Where(fmt.Sprintf("%s = ?", field), value).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
