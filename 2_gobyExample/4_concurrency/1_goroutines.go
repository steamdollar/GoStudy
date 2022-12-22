package main
import (
	"fmt"
	"time"
)

// execution의 경량 thread

func f(from string) {
	for i := 0; i <3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	// 이 일반적인 함수가 실행될때는 이 함수가 실행이 끝날때 까지 다음 코드를 실행하지 않음

	go f("goroutine")
	// 함수 f를 goroutine에 호출하려면 앞에 go를 붙이면 된다.
	// 이 새로운 goroutine은 호출자와 동시에 실행된다.
	// main함수와 f 함수가 동시에 실행됨

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// goroutine을 익명 함수 호출로도 시작할 수 있다.

	time.Sleep(time.Second)
	fmt.Println("done")
}