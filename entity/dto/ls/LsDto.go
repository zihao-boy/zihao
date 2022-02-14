package ls

type LsDto struct {
	GroupName string `json:"groupName"`
	Name      string `json:"name"`
	Privilege string `json:"privilege"`
	Size      int64  `json:"size"`
	LastModified string `json:"lastModified"` // Object last modified time
}
