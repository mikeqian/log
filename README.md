## what's the log for? ##

personal practice of golang and for future.

## base on ##

<pre>
https://github.com/qiniu/log
</pre>

## how to use? ##

It's very simple now.

<pre>
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
</pre>