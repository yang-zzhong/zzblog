package log

import (
	"fmt"
	l "log"
	"os"
	"sync"
	"time"
)

const (
	REQUEST          = "request"
	HTTP_PERFORMANCE = "http_performance"
	USER             = "user"
	SYSTEM           = "system"
	SPIDER           = "spider"
	DEBUG            = "debug"
)

var Path string
var pathfile string
var file *os.File
var std *l.Logger
var m sync.Mutex

func logger(filename string) *l.Logger {
	m.Lock()
	defer m.Unlock()
	var err error
	if Path == "" {
		Path = "/var/log/boo-blogger/"
	}
	if !exists(Path) {
		err = os.MkdirAll(Path, os.ModePerm)
	}
	if err != nil {
		panic(err)
	}
	should := Path + filename + "-" + time.Now().Format("2006-01-02") + ".log"
	if pathfile == should {
		return std
	}
	if file != nil {
		file.Close()
	}
	pathfile = should
	file, err = os.OpenFile(pathfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	std = l.New(file, "", l.LstdFlags)

	return std
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Print(filename string, v ...interface{}) {
	logger(filename).Output(2, fmt.Sprint(v...))
}

func Printf(filename string, format string, v ...interface{}) {
	logger(filename).Output(2, fmt.Sprintf(format, v...))
}
