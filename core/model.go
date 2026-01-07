package core

type Tool struct {
	Name string `yaml:"name"`
	Category string `yaml:"category"`
	Auth  string  `yaml:"auth"`
	Deployment string `yaml:"deployment"`
	Maintained bool `yaml:"maintained"`
}

