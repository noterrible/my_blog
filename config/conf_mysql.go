package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Config   string `yaml:"config"` //如charset的配置
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级，debug，打印sql。dev，production
}

func (m *Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Db + "?" + m.Config
}
