package main

import (
	"log"
	"os"
	"time"
)

func main() {
	name := time.Now().Format("20060102")

	f, err := os.OpenFile(name+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}
