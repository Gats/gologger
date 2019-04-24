package logger

import (
	"os"
	"log"
	"time"
	"fmt"
)

type Logger struct {
	Filename string
}

func New(file string) (*Logger, error) {
	l := new(Logger)
	l.Filename = file
	return l, nil
}

func fileOpen(file string) *os.File {
	f, err := os.OpenFile(file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("error opening file: %v", err)
	}
	
	return f
}

func (l *Logger) INFO(message string) {
	f := fileOpen(l.Filename)
	defer f.Close()
	log.SetOutput(f)
	log.SetPrefix("INFO: ")
    	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println(message)
}

func (l *Logger) ERROR(message string) {
	f := fileOpen(l.Filename)
	defer f.Close()
	log.SetOutput(f)
	log.SetPrefix("ERROR: ")
    	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println(message)
}

func (l *Logger) AppLogs(message, mtype string) {
	f := fileOpen(l.Filename)
	defer f.Close()
	now := time.Now()
	log.SetOutput(f)
	message = fmt.Sprintf("============================================== \n %s - %s \n %s \n", mtype, now, message)
	if _, err := f.WriteString(message); err != nil {
		log.Printf("%v", err)
	}

}


