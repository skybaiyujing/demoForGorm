package service

import (
	"Go_project/model"
	"gorm.io/gorm"
)

func DeleteUserByID(db *gorm.DB, id string) error {
	result := db.Delete(&model.User{}, "id = ?", id)
	return result.Error
}
func DeleteStudentBySno(db *gorm.DB, sno string) error {
	result := db.Delete(&model.Student{}, "sno = ?", sno)
	return result.Error
}
