package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	/* while(n < 5){
	   n++
	}
	*/
	for n < 5 {
		t.Log(n)
		n++
	}

	for n := 0; n < 100; n++ {
		t.Log(n)
	}
}

func TestInfiniteWhileLoop(t *testing.T) {
	n := 0
	for {
		if n == 100 {
			break
		}
		n++
	}
	t.Log(n)
}
