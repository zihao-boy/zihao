package menu

type MenuGroupDto struct {
	StatusCd    string `json:"statusCd"  sql:"-"`
	Seq         string `json:"seq"  `
	Name        string `json:"name" `
	Label       string `json:"label"`
	Icon        string `json:"icon"`
	GroupType   string `json:"groupType"  sql:"-"`
	Gid         string `json:"gId"  sql:"-"`
	Description string `json:"description" `
	CreateTime  string `json:"createTime"  sql:"-"`
}
