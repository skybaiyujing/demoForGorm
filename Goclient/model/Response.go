package model

/*
用于Get接口返回
*/
type StudentResponse struct {
	Data      []Student `json:"data"`       // 学生数据列表
	Total     int64     `json:"total"`      // 总记录数
	Page      int       `json:"page"`       // 当前页
	PageSize  int       `json:"page_size"`  // 每页记录数
	PageCount int64     `json:"page_count"` // 总页数
}
