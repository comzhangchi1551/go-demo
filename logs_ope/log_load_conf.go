package logs_ope

import (
	"encoding/json"
	"io"
	"os"
)

// 对应结构体
type LogConfig struct {
	LogDir   string `json:"log_dir"`
	LogLevel string `json:"log_level"`
}

func LoadLogConfig() *LogConfig {
	file, err := os.Open("confs/log_config.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	readAll, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var logConfig LogConfig
	err = json.Unmarshal(readAll, &logConfig)
	if err != nil {
		panic(err)
	}

	return &logConfig
}
