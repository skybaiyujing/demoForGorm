package service

import (
	"Go_project/model"
	"gorm.io/gorm"
	//"time"
)

/*
大写在数据库变小写，非首字母大写变成_
*/
func CreateTable(db *gorm.DB) error {
	//user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	err := db.AutoMigrate(&model.User{}) //结构体地址，要加{}
	if err != nil {
		//fmt.Println("创建表失败：", err)
		return err
	}
	//fmt.Println("创建表成功")
	return nil
}
