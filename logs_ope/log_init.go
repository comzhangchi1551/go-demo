package logs_ope

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logrus = logrus.New()

func init() {
	logConfig := LoadLogConfig()

	// 设置log文件，权限，是否自动创建等。
	file, err := os.OpenFile(logConfig.LogDir, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(err)
	}
	Logrus.Out = file

	// 设置日志级别
	log_level_map := map[string]logrus.Level{
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn":  logrus.WarnLevel,
		"info":  logrus.InfoLevel,
		"debug": logrus.DebugLevel,
		"trace": logrus.TraceLevel,
	}
	Logrus.SetLevel(log_level_map[logConfig.LogLevel])

	// 设置日志格式
	Logrus.SetFormatter(&logrus.TextFormatter{})
}
