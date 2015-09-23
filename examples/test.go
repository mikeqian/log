package main

import (
	"github.com/mikeqian/log"
	"time"
)

func main() {
	log.SetOutputLevel(log.Ldebug)

	start := time.Now()

	for i := 0; i < 1000; i++ {
		log.Debug("Debug: foo")
	}

	log.Debug(time.Since(start))
}
