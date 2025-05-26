package client

import (
	"GoClient/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func GetStuRaw(query model.StudentQuery) (model.StudentResponse, error) {
	url := "http://localhost:8080/students?" + query.ToQueryParams()
	resp, err := http.Get(url)
	if err != nil {
		return model.StudentResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var stuRes model.StudentResponse
	err = json.Unmarshal(body, &stuRes) //将 JSON 格式的字节数据（[]byte）反序列化成 Go中的结构体,即指针&stu指向的对象的类型。
	// 将结果赋值到该指针
	return stuRes, err
}
func GetStuRawBySno(sno int) (model.CommonResponse[[]model.Student], error) {
	url := "http://localhost:8080/students/" + strconv.Itoa(sno) // 避免直接字符串拼接
	resp, err := http.Get(url)
	if err != nil {
		return model.CommonResponse[[]model.Student]{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.CommonResponse[[]model.Student]{}, err
	}

	var stuRes model.CommonResponse[[]model.Student]
	err = json.Unmarshal(body, &stuRes) //将 JSON 格式的字节数据（[]byte）反序列化成 Go中的结构体,即指针&stu指向的对象的类型。
	// 将结果赋值到该指针
	return stuRes, err
}
func PostStuRaw(student model.Student) (*model.CommonResponse[model.Student], error) {
	url := "http://localhost:8080/students"

	jsonData, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var res model.CommonResponse[model.Student]
	err = json.Unmarshal(body, &res)
	return &res, err
}
