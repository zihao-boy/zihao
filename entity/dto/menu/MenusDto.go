package menu

type MenusDto struct {
	MId                  string `json:"mId" sql:"-"`
	MenuName             string `json:"menuName" sql:"-"`
	GId                  string `json:"gId" sql:"-"`
	Url                  string `json:"url" sql:"-"`
	MenuSeq              string `json:"menuSeq" sql:"-"`
	MenuDescription      string `json:"menuDescription" sql:"-"`
	MenuGroupName        string `json:"menuGroupName" sql:"-"`
	Icon                 string `json:"icon"`
	Label                string `json:"label"`
	MenuGroupSeq         string `json:"menuGroupSeq" sql:"-"`
	MenuGroupDescription string `json:"menuGroupDescription" sql:"-"`
	IsShow               string `json:"isShow" sql:"-"`
	Description          string `json:"description"`
}
