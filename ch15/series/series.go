package series

import (
	"errors"
	"fmt"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var largerThanHundredError = errors.New("n should be not lagre than 100")

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Square(n int) int {
	return n * n
}

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
