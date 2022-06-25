package logger

import (
	"testing"
)

func TestSnap(t *testing.T) {

	d := 10
	str := "Hello World"
	Debug("string")
	Debug("%v", d)
	Debug("%v", str)
	Debug("d = %v, str = %v", d, str)
	Error("string")
        Error("%v", d)
        Error("%v", str)
        Error("d = %v, str = %v", d, str)
}
