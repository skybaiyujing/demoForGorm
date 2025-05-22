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

##2.0版本：加入student结构体，其中有一个字段是结构体形式，名为FamilyInfo
```
type FamilyInfo struct {
	ID      uint `gorm:"primaryKey"`
	Sno     int  `gorm:"index"` // 外键
	Father  string
	Mother  string
	Address string
	Phone   string
}
type Student struct {
	Sno        int `gorm:"primaryKey"`
	Sname      string
	Ssex       string
	Sage       string
	FamilyInfo FamilyInfo `gorm:"foreignKey:Sno;references:Sno"`
}
```
##2.0版本：新加入姓名模糊搜索、性别筛选、年龄区间搜索、分页等功能
```
func QueryStudents(db *gorm.DB, query model.StudentQuery) ([]model.Student, int64, error) {
	var students []model.Student
	var total int64

	tx := db.Model(&model.Student{}).Preload("FamilyInfo")

	if query.Sname != "" {
		tx = tx.Where("sname LIKE ?", "%"+query.Sname+"%")
	}
	if query.Ssex != "" {
		tx = tx.Where("ssex = ?", query.Ssex)
	}
	if query.AgeMin != "" {
		tx = tx.Where("sage >= ?", query.AgeMin)
	}
	if query.AgeMax != "" {
		tx = tx.Where("sage <= ?", query.AgeMax)
	}
	tx.Count(&total)
	offset := (query.Page - 1) * query.PageSize
	result := tx.Offset(offset).Limit(query.PageSize).Find(&students)
	return students, total, result.Error
}
r.GET("/students", func(c *gin.Context) {
		query := model.StudentQuery{
			Sname:    c.Query("sname"), //c.GetQuery(key)
			Ssex:     c.Query("ssex"),
			AgeMin:   c.Query("age_min"),
			AgeMax:   c.Query("age_max"),
			Page:     utils.ParseIntDefault(c.Query("page"), 1),
			PageSize: utils.ParseIntDefault(c.Query("page_size"), 10),
		}
		students, total, err := service.QueryStudents(db, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data":       students,
			"total":      total,
			"page":       query.Page,
			"page_size":  query.PageSize,
			"page_count": (total + int64(query.PageSize) - 1) / int64(query.PageSize),
		})
	})
```
![image](https://github.com/user-attachments/assets/03684ddf-1577-4af4-8a0b-ffc0dd7cbaec)

##2.0版本删除事务
```
func DeleteStudentBySno(db *gorm.DB, sno string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("sno = ?", sno).Delete(&model.FamilyInfo{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.Student{}, "sno = ?", sno).Error; err != nil {
			return err
		}
		return nil
	})
}
```


##1.0版本：user结构体的GET、POST、DELETE接口
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
