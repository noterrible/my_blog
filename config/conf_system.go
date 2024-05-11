package config

type System struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	Env  string `yaml:"env"`
}

func (s *System) Addr() string {
	return s.Host + ":" + s.Port
}
