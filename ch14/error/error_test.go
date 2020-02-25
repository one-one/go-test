package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")

var largerThanHundredError = errors.New("n should be not lagre than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, largerThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(1000); err == nil {
		t.Log(v)
	} else {
		if err == LessThanTwoError {
			fmt.Println("num is less")
		} else if err == largerThanHundredError {
			fmt.Println("num is large")
		}

	}

}

func TestPaincExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("finally")
	//}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
}
