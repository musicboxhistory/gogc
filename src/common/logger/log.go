package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
)

func Init() {

}

func Snap(text string, value ...interface{}) {

	var print string

	pt, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pt).Name()
	print = filepath.Base(file)
	print = print + ":" + strconv.Itoa(line)
	print = print + "===>" + filepath.Base(funcName)
	print = print + "\t" + fmt.Sprintf(text, value...)
	log.Println(print)
}

func Error(text string, value ...interface{}) {

	var print string

	pt, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pt).Name()
	print = filepath.Base(file)
	print = print + ":" + strconv.Itoa(line)
	print = print + "===>" + filepath.Base(funcName)
	print = print + "\t" + fmt.Sprintf(text, value...)
	log.Println(print)
}
