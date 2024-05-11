package config

type Logger struct {
	Path         string `yaml:"path"`
	Level        string `yaml:"level"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      //显示日志行号
	LogInConsole bool   `yaml:"log_in_console"` //显示打印路径
}
