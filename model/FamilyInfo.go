package model

type FamilyInfo struct {
	ID      uint `gorm:"primaryKey"`
	Sno     int  `gorm:"index"` // 外键
	Father  string
	Mother  string
	Address string
	Phone   string
}
