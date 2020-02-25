package constant_test

import "testing"

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	t.Log(Readable, Writable, Executable)
}

func TestConstantTry1(t *testing.T) {
	a := 1 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
