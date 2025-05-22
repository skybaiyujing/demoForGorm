package model

type StudentQuery struct {
	Sname    string `form:"sname"`
	Ssex     string `form:"ssex"`
	AgeMin   string `form:"age_min"`
	AgeMax   string `form:"age_max"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}
