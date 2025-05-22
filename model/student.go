package model

// import "gorm.io/gorm"
type Student struct {
	Sno        int `gorm:"primaryKey"`
	Sname      string
	Ssex       string
	Sage       string
	FamilyInfo FamilyInfo `gorm:"foreignKey:Sno;references:Sno"`
}

//字段 结构体 家庭信息{123}
//列表查询 where 姓名模糊搜索、性别筛选、年龄区间筛选
//同时触发
//preload
//分页：页数、每页展示的个数
