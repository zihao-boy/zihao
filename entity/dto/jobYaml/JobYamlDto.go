package jobYaml

// job yaml
type JobYamlDto struct {
	Version string `yaml:"version"`
	Job JobDto `yaml:"job"`
}

type JobDto struct {
	JobName string `yaml:"job_name"`
	GitUrl string `yaml:"git_url"`
	GitPasswd string `yaml:"git_passwd"`
	GitUsername string `yaml:"git_username"`
	WorkDir string `yaml:"work_dir"`
	JobShell string `yaml:"job_shell"`
	Plans []JobPlanDto
}

type JobPlanDto struct {
	PackageUrl string `yaml:"package_url"`
	PackageName string `yaml:"package_name"`
	Path string `yaml:"path"`
	DockerfileName string `yaml:"dockerfile_name"`
	Dockerfile string `yaml:"dockerfile"`
	StartShell string `yaml:"start_shell"`
	ShellPath string `yaml:"shell_path"`
}
