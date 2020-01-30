package config

type Configuration struct {
	Syntax string `yaml:"syntax"validate:"required,syntax"`
	Tabs   []Tab  `yaml:"tabs,omitempty"validate:"omitempty,dive"`
}

type Tab struct {
	URL      string `yaml:"url"validate:"required"`
	Duration uint64 `yaml:"duration,omitempty"`
	Reload   bool   `yaml:"reload,omitempty"`
	Auth     Auth   `yaml:"auth,omitempty"validate:"omitempty,dive"`
	CSS      string `yaml:"css,omitempty"validate:"omitempty,uri"`
	JS       string `yaml:"js,omitempty"validate:"omitempty,uri"`
}

type Auth struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}
