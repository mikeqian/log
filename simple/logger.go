package main

import (
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "mike", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	l.Println("hello log")
}
