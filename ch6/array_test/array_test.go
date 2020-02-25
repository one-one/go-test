package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[2] = 10
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{2, 3, 4, 5, 6}
	t.Log(arr[1], arr[2], arr[0])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	arr3_sec := arr3[2:]
	t.Log(arr3_sec)
}

func TestMultiArrayInit(t *testing.T) {
	var a [3][2]int
	a = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	t.Log(a, a[1][1])

	var a2 = [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	t.Log(a2)
}
