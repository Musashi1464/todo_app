package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // ログの書き込み先を指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // ログのフォーマットを指定
	log.SetOutput(multiLogFile)                          // ログの出力先を指定
}
