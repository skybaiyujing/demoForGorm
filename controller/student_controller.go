package controller

import (
	"Go_project/model"
	"Go_project/service"
	"Go_project/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterStudentRoutes(r *gin.Engine, db *gorm.DB) {
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
	r.POST("/students", func(c *gin.Context) {
		var student model.Student
		//ShouldBindJSON（) 把客户端发来的 JSON 请求体绑定（解析）成你指定的结构体
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, model.ReturnType{
				Code:    400, //客户端错误
				Message: "参数错误: " + err.Error(),
				Data:    nil,
			})
			return
		}
		err := service.CreateStudentWithFamily(db, &student)
		//err := service.UpdateStudent(db, &student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ReturnType{
				Code:    500, //	服务端错误
				Message: "创建失败: " + err.Error(),
				Data:    nil,
			})
		}
		c.JSON(http.StatusCreated, model.ReturnType{
			Code:    200,
			Message: "创建成功",
			Data:    student,
		})
	})
	r.DELETE("/students/:sno", func(c *gin.Context) {
		sno := c.Param("sno") //“具体某个资源”，就用 c.Param() ；如果是“过滤多个资源”，才用 c.Query() 传查询参数。
		err := service.DeleteStudentBySno(db, sno)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
	})

}
