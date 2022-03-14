package service

type Page struct {
	Num  int `form:"page_num" json:"page_num"`   //頁數*筆數,從0(代表第一頁)開始
	Size int `form:"page_size" json:"page_size"` //從PageNum之後取出的筆數
}

func NewPage(num int, size int) *Page {
	return &Page{
		Num:  num,
		Size: size,
	}
}

func isValidPageNum(num int) bool {
	return num >= 1
}

func isValidPageSize(size int) bool {
	return size > 1
}

func (p *Page) GetSqlOffset() int {
	offset := 0
	if isValidPageNum(p.Num) {
		offset = (p.Num - 1) * p.Size
	}
	return offset
}

func (p *Page) GetSqlLimit() int {
	size := 1
	if isValidPageSize(p.Size) {
		size = p.Size
	}
	return size
}
