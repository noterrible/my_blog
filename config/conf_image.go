package config

type Image struct {
	Path         string `yaml:"path"`
	enableOrigin string `yaml:"enable_origin"`
	Size         int64  `yaml:"size"`
}
