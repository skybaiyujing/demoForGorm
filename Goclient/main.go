package main

import (
	"GoClient/client"
	"GoClient/model"
	"fmt"
)

func main() {
	//query := model.StudentQuery{
	//	Sname:    "",
	//	Ssex:     "",
	//	AgeMin:   "18",
	//	AgeMax:   "25",
	//	Page:     1,
	//	PageSize: 10,
	//}

	student := model.Student{
		Sno:   1023,
		Sname: "张宇",
		Ssex:  "女",
		Sage:  "30",
		FamilyInfo: model.FamilyInfo{
			Father:  "张爸爸",
			Mother:  "张妈妈",
			Address: "北京市朝阳区xx路",
			Phone:   "13800001111",
		},
	}

	// 使用 net/http 测试
	//students, err := client.GetStuRaw(query)
	//fmt.Println("Students from net/http:", students, "ERR:", err)

	//res1, err := client.PostStuRaw(student)
	//fmt.Println("POST:", res1, "ERR:", err)

	// 使用 resty 测试
	fmt.Println("=== [resty] ===")
	//
	//students2, err := client.GetStudentsResty(query)
	//fmt.Println("GET:", students2, "ERR:", err)

	res2, err := client.PostStudentResty(student)
	fmt.Println("POST:", res2, "ERR:", err)
}
