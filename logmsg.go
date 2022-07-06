package swordtech

import (
	"fmt"
	"github.com/markoxley/daggertech"
	"net/http"
)

var codes = map[int]string{
	-1: "System Fail",
	0:  "System Success",
}

// LogMsg is the log message to be stored in the database
type LogMsg struct {
	daggertech.Model
	IP                string `daggertech:"size:20"`
	Method            string `daggertech:"size:32"`
	Text              string `daggertech:""`
	Result            int    `daggertech:""`
	ResultDescription string `daggertech:""`
}

// ToString returns the string representation of the error
func (l *LogMsg) ToString() string {
	return fmt.Sprintf("%s\t%s\t%v\t[%d]\t%s", l.IP, l.Method, l.Text, l.Result, l.ResultDescription)
}

// newLogMsg creates a new LogMsg model and saves it to the database
func newLogMsg(method string, ip, text string, result int) *LogMsg {
	description, ok := codes[result]
	if !ok {
		if description = http.StatusText(result); description == "" {
			description = "Unknown Response"
		}
	}
	logMsg := &LogMsg{
		Method:            method,
		IP:                ip,
		Text:              text,
		Result:            result,
		ResultDescription: description,
	}
	daggertech.Save(logMsg)
	return logMsg
}
