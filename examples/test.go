package main

import (
	"github.com/mikeqian/log"
	"os"
	"time"
)

func main() {
	log.InitLogger(os.Stdout)

	start := time.Now()
	for i := 0; i < 5; i++ {
		log.Debug("1")
		log.Info("2")
		log.Error("3")
	}

	log.Info(time.Since(start))
}
