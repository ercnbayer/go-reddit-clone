package logger

import (
	"log"
)

var LogLevel uint8

func Fatal(v ...interface{}) {

	if LogLevel > 0 {
		log.Fatal(v...)
	}

}

func Info(v ...interface{}) {

	if LogLevel > 1 {
		log.Println(v...)
	}

}

func Error(v ...interface{}) {

	if LogLevel > 2 {
		log.Println(v...)
	}

}
