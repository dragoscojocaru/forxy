package logger

import (
	"log"
	"os"
)

func FileErrorLog(err error) {
	f, errFileOpen := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errFileOpen != nil {
		log.Fatalf("error opening log file: %v", errFileOpen)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(err.Error())
}
