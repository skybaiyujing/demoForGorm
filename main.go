package main

import (
	"Go_project/controller"
	"Go_project/db"
	"Go_project/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.InitDB()
	service.CreateTable(db)
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	r := gin.Default()
	// 注册用户相关接口
	controller.RegisterUserRoutes(r, db)
	r.Run(":8080")

	//service.Insert(db)
	//{
	//	var user model.User
	//	result := db.First(&user, 1)
	//	if result.Error != nil {
	//		log.Println("查询失败:", result.Error)
	//		return
	//	}
	//	user.Name = "hhh"
	//	user.Age = 101
	//	service.UpdateUser(db, &user)
	//	//db.Save(&user)
	//}
	//user, err := service.SelectWithId(db)
	//if err != nil {
	//	log.Println("查询失败:", err)
	//	return
	//}
	//for i := 0; i < len(user); i++ {
	//	fmt.Printf("姓名: %s, 年龄: %d\n", user[i].Name, user[i].Age)
	//}
}
