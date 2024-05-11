package initialization

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"my_blog/global"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//颜色设置
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("20060102 15:04:05")
	if entry.HasCaller() {
		//自定义路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s \n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                        //新建实例
	mLog.SetOutput(os.Stdout)                                   //设置输出类型为命令行
	mLog.SetReportCaller(global.Config.Logger.ShowLine)         //是否显示函数名行号
	mLog.SetFormatter(&LogFormatter{})                          //自定义日志格式
	level, err := logrus.ParseLevel(global.Config.Logger.Level) //设置日志等级
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //日志等级
	InitDefaultLogger()
	return mLog
}
func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(global.Config.Logger.Level) //设置日志等级
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //日志等级
}
