package csp

import (
	"fmt"
	"testing"
	"time"
)

func servcie() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(servcie())
	otherTask()
}

func AsyncServic() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := servcie()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {

	retCh := AsyncServic()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(1 * time.Second)
}
