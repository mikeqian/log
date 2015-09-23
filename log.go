package log

import (
	"fmt"
	"io"
	"os"

	"sync"
	"time"
)

var (
	stdFlags  = lstdFlags | lshortfile
	loggerStd = &logger{out: os.Stderr, flag: stdFlags}
	recvBytes chan []byte
	recvOver  = make(chan bool)
)

var (
	Debug logType
	Info  logType
	Error logType
)

const (
	infoStr  = "[Info ]- "
	debugStr = "[Debug]- "
	errorStr = "[Error]- "
)
const (
	LDebug = 1 << iota
	LInfo
	LError
	LFatal
)
const (
	_debug = 1<<(iota+1) - 1
	_info
	_error
	_fatal
)
const (
	ldate = 1 << iota
	ltime
	lmicroseconds
	llongfile
	lshortfile
	lstdFlags = ldate | ltime
)

type logfType func(format string, v ...interface{})

type logType func(v ...interface{})

func Close() {
	close(recvBytes)
	<-recvOver
}

func InitLogger(w io.Writer) {
	Debug = makeLog(debugStr)
	Info = makeLog(infoStr)
	Error = makeLog(errorStr)
}

func makeLog(prefix string) (y logType) {
	return func(v ...interface{}) {
		loggerStd.write(prefix, fmt.Sprintln(v...))
	}
}

type logger struct {
	mutex sync.Mutex
	flag  int
	out   io.Writer
	buf   []byte
}

func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}

func (l *logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	if l.flag&(ldate|ltime|lmicroseconds) != 0 {
		if l.flag&ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(ltime|lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(lshortfile|llongfile) != 0 {
		if l.flag&lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

func (l *logger) write(prefix string, s string) error {
	now := time.Now()
	var file string
	var line int
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buf = l.buf[:0]
	l.buf = append(l.buf, prefix...)

	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) > 0 && s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	newSlice := make([]byte, len(l.buf))
	if recvBytes != nil {
		copy(newSlice, l.buf)
		recvBytes <- newSlice
	}
	return err
}
