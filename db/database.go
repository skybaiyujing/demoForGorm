package db

import (
	//"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open("sqlserver://sa:1234qwer@192.168.246.183:1433?database=ST"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
