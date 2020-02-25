package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `key:"value"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookID string
	Name   string
	Age    int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name field.'")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("key"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
