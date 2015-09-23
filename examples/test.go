package main

import (
	"github.com/mikeqian/log"
	"time"
)

func main() {
	log.SetOutputLevel(Ldebug)

	start := time.Now()

	for i := 0; i < 5; i++ {
		Debug("Debug: foo")
	}

	log.Debug(time.Since(start))
}
