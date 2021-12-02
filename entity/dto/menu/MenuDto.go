package menu

type MenuDto struct {
	Url         string `json:"url" `
	StatusCd    string `json:"statusCd"  sql:"-"`
	Seq         string `json:"seq"  `
	Name        string `json:"name" `
	Mid         string `json:"mId"  sql:"-"`
	IsShow      string `json:"isShow"  sql:"-"`
	Gid         string `json:"gId"  sql:"-"`
	Description string `json:"description" `
	CreateTime  string `json:"createTime"  sql:"-"`
}
