package util

import (
	"encoding/json"
	"fmt"
	"time"
)

type LogLevel string

var (
	LOG_LEVEL_INFO  = LogLevel("INFO")
	LOG_LEVEL_ERROR = LogLevel("ERROR")
)

func Log(level LogLevel, tag string, context interface{}) {
	ctxStr := ""
	bytes, err := json.Marshal(context)
	if err != nil {
		ctxStr = ""
	} else {
		ctxStr = string(bytes)
	}

	fmt.Println(fmt.Sprintf("%s	%s	%s	%s",
		time.Now().Format("2006-01-02 15:04:05"), level, tag, ctxStr))
}

func InfoLog(tag string, context interface{}) {
	Log(LOG_LEVEL_INFO, tag, context)
}

func ErrorLog(tag string, context interface{}) {
	Log(LOG_LEVEL_ERROR, tag, context)
}
