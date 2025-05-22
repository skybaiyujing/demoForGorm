package service

import (
	"Go_project/model"
	"errors"
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
func CreateStudentWithFamily(db *gorm.DB, student *model.Student) error {
	// 防御性编程：判断 student.Sno 和 FamilyInfo.Sno 是否一致
	if student.Sno == 0 {
		return errors.New("学生 Sno 不能为空")
	}
	student.FamilyInfo.Sno = student.Sno
	// 使用事务确保原子性：要么都成功，要么都失败
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&student).Error; err != nil {
			return err
		}
		return nil
	})
}
