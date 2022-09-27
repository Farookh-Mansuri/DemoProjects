package main

import (
	"log"
	"os"
)

func main() {
	logFile := openLogFile()
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("This is logging!")
}
func openLogFile() *os.File {
	log.Println("This is open!")
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
