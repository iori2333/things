package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var (
	loggers map[string]*log.Logger
	once    sync.Once
	lock    sync.Mutex
)

const FLAG = log.LstdFlags | log.Lshortfile

var Logger = GetLogger("SYSTEM", os.Stderr)

func GetLogger(name string, out io.Writer) *log.Logger {
	once.Do(func() {
		loggers = make(map[string]*log.Logger)
	})
	lock.Lock()
	defer lock.Unlock()
	if _, ok := loggers[name]; !ok {
		loggers[name] = log.New(out, fmt.Sprintf("[%s] ", name), FLAG)
	}
	return loggers[name]
}
