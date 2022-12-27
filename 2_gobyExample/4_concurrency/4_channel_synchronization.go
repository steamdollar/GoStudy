package main
import (
	"fmt"
	"time"
)

/*
	channel을 이용해 goroutine들의 실행을 동기화 할 수 있다.
	예시로 고루틴의 종료를 기다리기 위해 blocking receive를 사용한다.

	여러 개의 고루틴이 종료 되는 것을 기다릴 때는 waitgroup을 상요하는 것이 좋다.
*/

func worker(done chan bool) {
	// 고루틴에서 돌릴 함수
	// done channel은 다른 고루틴에 이 함수의 작동이 끝났다는 것을 알려주기 위해 사용된다.

	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	// 할일을 다 헀으면 value를 전송해 종료를 notify
}

func main() {
	done := make(chan bool, 1)
	// channel의 nofity 확인 후, 고루틴을 시작
	go worker(done)

	<- done
	// 채널에서 값을 전달 받을때까지 여기서 block 되어있다.
	// 이거 없으면 worker 시작전에 main 함수가 끝나고 프로그램이 종료되서
	// worker 함수는 실행되지도 못함
}