package logs

import (
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func Init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	// log.SetOutput(mw)
}
func Debug(args ...interface{}) {
	log.Printf("\nWARNING : %v \ttime %v", args, time.Now())
	WarningLogger.Println(args...)
}
func Error(args ...interface{}) {
	log.Printf("\nERROR : %v \ttime %v", args, time.Now())
	ErrorLogger.Println(args...)
}
func Info(args ...interface{}) {
	log.Printf("\nINFO : %v \ttime %v", args, time.Now())
	InfoLogger.Println(args...)
}
