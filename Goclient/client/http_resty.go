package client

import (
	"GoClient/model"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
)

var client = resty.New()

func GetStudentsResty(query model.StudentQuery) (model.StudentResponse, error) {
	var response model.StudentResponse
	_, err := client.R().
		SetQueryParamsFromValues(url.Values{
			"sname":     []string{query.Sname},
			"ssex":      []string{query.Ssex},
			"age_min":   []string{query.AgeMin},
			"age_max":   []string{query.AgeMax},
			"page":      []string{strconv.Itoa(query.Page)},
			"page_size": []string{strconv.Itoa(query.PageSize)},
		}).
		SetResult(&response). //SetResult设置响应的目标对象类型
		Get("http://localhost:8080/students")

	if err != nil {
		return model.StudentResponse{}, err
	}
	return response, nil
}
func PostStudentResty(student model.Student) (model.CommonResponse[model.Student], error) {
	// 创建一个新的 StudentResponse 来接收数据
	var response model.CommonResponse[model.Student]

	// 发送 POST 请求
	//resp是请求返回的原始响应对象，包含了所有关于 HTTP 请求的元数据
	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(student).     // 设置请求体为 student
		SetResult(&response). // 设置解析结果为 StudentResponse 类型
		Post("http://localhost:8080/students")
	// 错误处理
	if err != nil {
		return model.CommonResponse[model.Student]{}, err
	}

	// 返回解析后的响应数据
	return response, nil
}
