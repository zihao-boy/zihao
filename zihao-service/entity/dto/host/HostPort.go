package host

type HostPortDto struct {
	Protocol string `json:"protocol"`
	Port string `json:"port"`
	ProgramName string `json:"programName"`
}
