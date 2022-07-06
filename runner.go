package swordtech

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var batchProcess []func() error
var backgroundProcess []func() error

func init() {
	batchProcess = make([]func() error, 0)
	backgroundProcess = make([]func() error, 0)
}

// AddBatchProcess adds a function to be executed once a day
func AddBatchProcess(f func() error) {
	batchProcess = append(batchProcess, f)
}

// AddBackgroundProcess adds a function that will be executed in the background
func AddBackgroundProcess(f func() error) {
	backgroundProcess = append(backgroundProcess, f)
}

func runProcesses() {
	backgroundIndex := 0
	batchIndex := 0
	inBatch := false
	interval := 30
	for {
		delay := time.NewTimer(time.Second * time.Duration(interval))
		<-delay.C
		interval = int(GetParameter("Batch", "Interval", "The interval between runner executions", "60").Int())
		batchHour := int(GetParameter("Batch", "Hour", "Hour to execite batch processes", "2").Int())

		if !inBatch {
			now := time.Now()
			lastBatch := GetParameter("Batch", "LastBatch", "Time of the last batch process run", "")
			if lastBatch.Time() == nil {
				inBatch = true
			} else {
				next := lastBatch.Time().Add(time.Hour * 24)
				if now.After(next) {
					inBatch = true
				}
			}
			if inBatch {
				u := time.Date(now.Year(), now.Month(), now.Day()+1, batchHour, 0, 0, 0, time.UTC)
				lastBatch.Set(u)
				batchIndex = 0
			}
		}

		if inBatch {
			if batchIndex >= len(batchProcess) {
				inBatch = false
				continue
			}
			runProcess(batchProcess[batchIndex])
			batchIndex++
		} else {
			if backgroundIndex >= len(backgroundProcess) {
				backgroundIndex = 0
				continue
			}
			runProcess(backgroundProcess[backgroundIndex])
			backgroundIndex++
		}
	}

}

func runProcess(p func() error) {
	fullname := strings.Split(runtime.FuncForPC(reflect.ValueOf(p).Pointer()).Name(), ".")
	name := fullname[len(fullname)-1]
	defer func() {
		if x := recover(); x != nil {
			Log(nil, nil, fmt.Sprintf("Panic running %s : %v", name, x), -1)
		} else {
			Log(nil, nil, fmt.Sprintf("%s successfully completed", name), 0)
		}
	}()
	if err := p(); err != nil {
		Log(nil, nil, fmt.Sprintf("Error running %s : %v", name, err), -1)
	}
}
