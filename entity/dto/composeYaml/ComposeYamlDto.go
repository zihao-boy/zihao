package composeYaml

// docker compose 
type ComposeYamlDto struct {
	Version string `yaml:"version"` // docker compose version
	Services []interface{} `yaml:"services"`
}

type ComposeYamlZiHaoDto struct {
	Version string `yaml:"version"` // docker compose version
	Services []interface{} `yaml:"services"`
	ZihaoCmd string `yaml:"zihao_cmd" json:"zihao_cmd"`
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
	Image string `yaml:"image" json:"image"`
	Volumes []string `yaml:"volumes" json:"volumes"`
	Ports []string `yaml:"ports" json:"ports"`
	ExtraHosts []string `yaml:"extra_hosts" json:"extra_hosts"`
	Environment []string `yaml:"environment" json:"environment"`
}