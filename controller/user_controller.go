package controller

import (
	"Go_project/model"
	"Go_project/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", func(c *gin.Context) {
		users, err := service.SelectAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
	r.GET("/users/", func(c *gin.Context) {
		field := c.Query("field") // 比如 "id" 或 "name"
		value := c.Query("value") // 比如 "123" 或 "Tony"

		var users []model.User
		var err error

		switch field {
		case "id":
			users, err = service.SelectWithField(db, "id", value)
		case "name":
			users, err = service.SelectWithField(db, "name", value)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的查询字段"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(users) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "未找到用户"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.UpdateUser(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service.DeleteUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
	})
}
