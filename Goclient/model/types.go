package model

type FamilyInfo struct {
	ID      uint   `json:"id"`
	Sno     int    `json:"sno"`
	Father  string `json:"father"`
	Mother  string `json:"mother"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Student struct {
	Sno        int        `json:"sno"`
	Sname      string     `json:"sname"`
	Ssex       string     `json:"ssex"`
	Sage       string     `json:"sage"`
	FamilyInfo FamilyInfo `gorm:"foreignKey:Sno;references:Sno"`
}

/*
用于POST返回
*/
type CommonResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"` //泛型，调用时CommonResponse[Student]
}
