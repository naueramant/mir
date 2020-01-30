package config

type Configuration struct {
	Syntax string `yaml:"syntax"validate:"required,syntax"`
	Tabs   []Tab  `yaml:"tabs,omitempty"validate:"omitempty,dive"`
	Jobs   []Job  `yaml:"jobs,omitempty"validate:"omitempty,dive"`
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

type Job struct {
	Type string `yaml:"type"validate:"required"`
	When string `yaml:"when"validate:"required"`
	Data Data   `yaml:"data,omitempty"validate:"omitempty,dive"`
}

type Data struct {
	// Command
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`

	// Tab, Message
	Duration uint64 `yaml:"duration"`

	// Tab
	URL string `yaml:"url"`

	// Message
	Message         string `yaml:"message"`
	FontSize        uint64 `yaml:"fontSize"`
	TextColor       string `yaml:"textColor"`
	BackgroundColor string `yaml:"backgroundColor"`
	Blink           bool   `yaml:"blink"`
}
