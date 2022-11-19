package model

// BasePage 分页操作
type BasePage struct {
	PageNum  int `form:"page_num"`  //第几页,即mysql中的offset
	PageSize int `form:"page_size"` //每一页显示的个数,即mysql中的limit
}
