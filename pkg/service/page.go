package service

type Page struct {
	Num  int `form:"page_num" json:"page_num"`   //頁數*筆數,從0(代表第一頁)開始
	Size int `form:"page_size" json:"page_size"` //從PageNum之後取出的筆數
}
