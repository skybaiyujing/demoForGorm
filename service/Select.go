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

func SelectStuWithSno(db *gorm.DB, ID string) ([]model.Student, error) {
	//根据主键检索
	var stu []model.Student
	result := db.Where("sno = ?", ID).Preload("FamilyInfo").Find(&stu)
	if result.Error != nil {
		return nil, result.Error
	}
	return stu, nil
}
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
