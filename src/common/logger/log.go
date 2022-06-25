package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func Init() {

}

func Debug(text string, value ...interface{}) {

	var print string

	if os.Getenv("DEBUG_MODE") == "OFF" {
		return
	}

	pt, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pt).Name()
	print = "[DEBUG]"
	print = print + filepath.Base(file)
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
	print = "[ERROR]"
	print = print + filepath.Base(file)
	print = print + ":" + strconv.Itoa(line)
	print = print + "===>" + filepath.Base(funcName)
	print = print + "\t" + fmt.Sprintf(text, value...)
	log.Println(print)
}
