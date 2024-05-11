package config

type SiteInfo struct {
	Title    string `yaml:"title"`
	Job      string `yaml:"job"`
	Addr     string `yaml:"addr"`
	CreateAt string `yaml:"create_at"`
	Github   string `yaml:"github"`
}
