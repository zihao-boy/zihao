package host

type HostDiskDto struct {
	DiskName string `json:"diskName"`
	Size     string `json:"size"`
	FreeSize string `json:"freeSize"`
	Dir      string `json:"dir"`
}
