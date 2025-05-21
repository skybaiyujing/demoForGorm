package service

import (
	"Go_project/model"
	"gorm.io/gorm"
	"time"
	//"time"
)

func Insert(db *gorm.DB) *gorm.DB {
	users := []*model.User{
		{Name: "Tony", Age: 20, Birthday: time.Now(), Email: "227@new"},
		{Name: "Jenny", Age: 20, Birthday: time.Now(), Email: "217@new"},
	}
	result := db.Create(&users)
	return result
}
func CreateSelect(db *gorm.DB) *gorm.DB {
	//user := model.User{Name: "Tony", Age: 20, Birthday: time.Now(), Email: "227@new"}
	user2 := model.User{Name: "Jenny", Age: 20, Birthday: time.Now(), Email: "217@new"}
	result := db.Select("Name").Create(&user2)
	return result
}
