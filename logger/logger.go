package logger

import (
	"log"
)

type LogLevels int64

const (
	OnlyFatalLogs = iota + 1
	ErrorAndFatalLogs
	AllLogs
)

var LogLevel LogLevels

func Fatal(v ...interface{}) { //log fatal errors

	if LogLevel > 0 {
		log.Fatal(v...)
	}

}

func Info(v ...interface{}) { // log infos

	if LogLevel > 1 {
		log.Println(v...)
	}

}

func Error(v ...interface{}) { // log errors

	if LogLevel > 2 {
		log.Println(v...)
	}

}
