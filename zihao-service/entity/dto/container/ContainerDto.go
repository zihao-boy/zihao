package container

type ContainerDto struct {
	ContainerId string `json:"containerId"`
	ContainerName string `json:"containerName"`
	Image string `json:"image"`
	Port string `json:"port"`
}
