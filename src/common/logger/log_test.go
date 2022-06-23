package logger

import (
	"testing"
)

func TestSnap(t *testing.T) {

	d := 10
	str := "Hello World"
	Snap("string")
	Snap("%v", d)
	Snap("%v", str)
	Snap("d = %v, str = %v", d, str)
}
