package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_Level"` //日志等级，debug，打印sql。dev，production
}
type Logger struct {
	Level        string `yaml:"level"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_Line"`      //显示日志行号
	LogInConsole bool   `yaml:"log_In_Console"` //显示打印路径
}
type System struct {
	Port string `yaml:"port"`
	Host int    `yaml:"host"`
	Env  string `yaml:"env"`
}
