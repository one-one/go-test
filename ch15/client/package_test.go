package client

import "testing"
import "goTest/demo/ch15/series"

func TestPackage(t *testing.T) {

	if value, err := series.GetFibonacci(-5); err == nil {
		t.Log(value)
	} else {
		t.Log(err)
	}

	t.Log(series.Square(5))

}
