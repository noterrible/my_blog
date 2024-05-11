package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	JWT      JWT      `yaml:"jwt"`
	Image    Image    `yaml:"image"`
	SiteInfo SiteInfo `yaml:"site_info"`
}
