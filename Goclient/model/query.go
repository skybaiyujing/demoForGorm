package model

import (
	"net/url"
	"strconv"
)

type StudentQuery struct { //用于GET语句时的筛选、分页
	Sname    string `form:"sname"`
	Ssex     string `form:"ssex"`
	AgeMin   string `form:"age_min"`
	AgeMax   string `form:"age_max"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

func (q *StudentQuery) ToQueryParams() string {
	values := url.Values{}
	//一个 map，key 是字符串，value 是字符串数组
	if q.Sname != "" {
		values.Set("sname", q.Sname)
	}
	if q.Ssex != "" {
		values.Set("ssex", q.Ssex)
	}
	if q.AgeMin != "" {
		values.Set("age_min", q.AgeMin)
	}
	if q.AgeMax != "" {
		values.Set("age_max", q.AgeMax)
	}
	if q.Page > 0 {
		values.Set("page", strconv.Itoa(q.Page))
	}
	if q.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(q.PageSize))
	}

	return values.Encode()
	//把 map 转成 URL 查询字符串
}
