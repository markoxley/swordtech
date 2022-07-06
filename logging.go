package swordtech

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var showLog bool

func writeLog(method *string, ip *string, text string, result int, panic bool) {
	outputIP := ""
	if ip != nil {
		outputIP = *ip
	}

	outputMethod := "Internal"
	if method != nil {
		outputMethod = *method
	}

	re := regexp.MustCompile(`(?m)\[([^\]]*)\]`)

	if len(re.FindStringIndex(outputIP)) > 0 {
		outputIP = re.FindStringSubmatch(outputIP)[1]
	}

	if idx := strings.Index(outputIP, ":"); idx > -1 {
		outputIP = outputIP[0:idx]
	}

	logMsg := newLogMsg(outputMethod, outputIP, text, result)
	logfile := getLogPath() + strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", "") + ".log"
	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	log.SetOutput(file)
	out := logMsg.ToString()
	log.Println(out)

	if showLog {
		if !panic {
			fmt.Println(out)
		} else {
			log.Panic(out)
		}
	}
}

func getLogPath() string {
	logPath := "logs"
	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		if err = os.Mkdir(logPath, os.ModeDir|os.ModePerm); err != nil {
			panic("Unable to create log folder")
		}
	}
	return logPath + "/log_"
}

// Log outputs a message to the daily log file
func Log(method *string, ip *string, text string, result int) {
	writeLog(method, ip, text, result, false)
}

// Panic logs the entry and raises a panic
func Panic(method *string, ip *string, text string, result int) {
	writeLog(method, ip, text, result, true)
}
