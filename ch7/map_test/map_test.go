package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessIsExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Log(v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestMapTravel(t *testing.T) {
	//m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	var m2 map[int]int
	m3 := map[int]int{}
	//m3[4] = 16
	if m2 == nil {
		t.Log("m2 is nil")
	}
	for k, v := range m3 {
		t.Log(k, v)
	}
}

func TestMapWithFunc(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1

	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)

	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)

	}
	t.Log(len(mySet))

}
