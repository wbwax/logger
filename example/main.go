package main

import (
	"github.com/wbwax/logger"
)

func main() {
	config := logger.Config{
		MaxSize:    1,         // 1 MB
		MaxAge:     1,         // 1 day
		MaxBackups: 10,        // 10
		Level:      "info",    // logger level
		Path:       "logs",    // path
		Encoding:   "console", // only support "console" and "json"
	}
	logger.Init(config)
	defer logger.Sync() // flush buffer, if any
	logger.Infof("msg=%s||level=%s", "succeed to init logger", config.Level)
	logger.Debugf("msg=debug log") // no debug log file in logs directory
	logger.Infof("msg=info log")
	logger.Warnf("msg=warn log")
	logger.Errorf("msg=error log")
}
