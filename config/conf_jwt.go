package config

type JWT struct {
	Secret  string `yaml:"secret"`
	Expires int64  `yaml:"expires"`
}
