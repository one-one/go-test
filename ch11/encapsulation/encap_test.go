package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 22}
	e1 := Employee{Name: "mike", Age: 23}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 25

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {

	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) Stringss() string {
	fmt.Printf("access is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 22}
	fmt.Printf("access is %x", unsafe.Pointer(&e.Name))
	t.Log(e.Stringss())

}
