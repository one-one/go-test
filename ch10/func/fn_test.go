package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func fnMapTest(v map[int]int) int {
	v[2] = 30
	return v[2]
}

func TestFn(t *testing.T) {
	var aa map[int]int = make(map[int]int, 10)
	aa[1] = 10
	aa[2] = 20
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
	t.Log(aa[2])
	t.Log(fnMapTest(aa))
	t.Log(aa[2])

}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 10)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}

	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start11")
	panic("Fatal error")
	fmt.Println("End")
	t.Log("end")
}
