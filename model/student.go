package model

// import "gorm.io/gorm"
type Student struct {
	Sno   string `gorm:"primaryKey"`
	Sname string
	Ssex  string
	Sage  string
	ASD   string
}
