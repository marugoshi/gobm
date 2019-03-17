package app_log

import (
	"io"
	"log"
	"os"
)

var (
	debug = log.New(os.Stderr, "[DEBUG] ", log.LstdFlags)
	info  = log.New(os.Stderr, "[INFO] ", log.LstdFlags)
	warn  = log.New(os.Stderr, "[WARN] ", log.LstdFlags)
	fatal = log.New(os.Stderr, "[FATAL] ", log.LstdFlags)
)

func Debugf(fmt string, v ...interface{}) {
	debug.Printf(fmt, v...)
}

func Debug(msg interface{}) {
	debug.Println(msg)
}

func Infof(fmt string, v ...interface{}) {
	info.Printf(fmt, v...)
}

func Info(msg interface{}) {
	info.Println(msg)
}

func Warnf(fmt string, v ...interface{}) {
	warn.Printf(fmt, v...)
}

func Warn(msg interface{}) {
	warn.Println(msg)
}

func Fatalf(fmt string, v ...interface{}) {
	fatal.Printf(fmt, v...)
}

func Fatal(msg interface{}) {
	fatal.Println(msg)
}

func SetOutput(w io.Writer) {
	debug.SetOutput(w)
	info.SetOutput(w)
	warn.SetOutput(w)
	fatal.SetOutput(w)
}
