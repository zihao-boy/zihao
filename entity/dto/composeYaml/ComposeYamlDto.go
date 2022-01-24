package composeYaml

// docker compose 
type ComposeYamlDto struct {

	Version string `yaml:"version"` // docker compose version
	Services []interface{} `yaml:"services"`
}

type ServicesDto struct {
	Name string
	Service ServiceDto
}

func (n *ServicesDto) ToMap() interface{} {
	yamlMap := make(map[string]interface{})
	yamlMap[n.Name] = n.Service
	return yamlMap
}

type ServiceDto struct {
	Image string `yaml:"image"`
	Volumes []string `yaml:"volumes"`
	Ports []string `yaml:"ports"`
	ExtraHosts []string `yaml:"extra_hosts"`
	Environment []string `yaml:"environment"`
}