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
