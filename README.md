# demoForGorm
##建立数据库连接
```
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:1234qwer@192.168.246.183:1433?database=ST"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
```
##GET、POST、DELETE接口
- `GET /users`：获取所有用户
```
	r.GET("/users", func(c *gin.Context) {
		users, err := service.SelectAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
```
- `GET /users/?field={id|name}&value={值}`：按字段查询用户
```
func SelectWithField(db *gorm.DB, field string, value interface{}) ([]model.User, error) {
	var users []model.User
	result := db.Where(fmt.Sprintf("%s = ?", field), value).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
```

```
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
```
- `POST /users`：添加或更新用户（JSON 请求体）
```
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
```
- `DELETE /users/:id`：按 ID 删除用户
```
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service.DeleteUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
	})
```
