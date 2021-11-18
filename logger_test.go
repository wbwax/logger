package logger

import (
	"fmt"
	"testing"
)

var (
	config Config
)

func init() {
	config = Config{
		MaxSize:    1,         // 1 MB
		MaxAge:     1,         // 1 day
		MaxBackups: 10,        // 10
		Level:      "info",    // logger level
		Path:       "logs",    // path
		Encoding:   "console", // only support "console" and "json"
	}
}

func TestInit(t *testing.T) {
	// init logger
	if err := Init(config); err != nil {
		t.Errorf("msg=%s||err=%s", "failed to init logger", err.Error())
		return
	}
	fmt.Println("succeed to init logger")

	// test logger
	Infof("msg=succeed to init logger||num=%d", 110)
}
