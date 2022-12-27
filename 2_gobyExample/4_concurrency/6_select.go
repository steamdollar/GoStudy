package main
import (
	"fmt"
	"time"
)

/*
	select는 여러 개의 채널의 작동을 기다릴 수 있게 해준다.
	고루틴과 채널을 select로 묶어줄 수 있다는게 go의 장점이라카더라.
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	// 채널 2개를 만듦

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// 각 채널은 정해진 시간 후 value를 받는다.

	for i := 0; i <2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// 반복문, select를 이용해 각 value를 동시에 기다리고 받는 순서대로 출력한다.
}
