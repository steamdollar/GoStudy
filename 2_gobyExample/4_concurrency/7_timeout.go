package main
import (
	"fmt"
	"time"
)

/*
	timeout은 외부 리소스와 연결되거나 실행 시간을 bind해야 하는 프로그램에 중요하다.
	go에서 cahnnel과 select를 이용해 timeout을 쉽게 구현할 수 있다.
*/

func main() {
	c1 := make(chan string, 1)
	go func () {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// chan c1에 2s 후 데이터를 전송하는 함수가 있다.
	// 이 때 채널은 buffered channel이라 고루틴은 block되지 않는다.
	// 이러면 채널이 읽히지 않아 프로그램 전체가 멈추는 사태를 일어나지 않게 해줄 수 있다.

	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	// 이게 select를 이용해 timeout을 구현한 코드이다.
	// <- c1은 결과를 기다리고,
	// <- time.After는 1초간 기다린 후, 뭐가 오지 않으면 (다른 case가 실행되지 않으면)
	// timeout을 출력한다.

	c2 := make(chan string, 1)
	go func () {
		time.Sleep( 2 * time.Second)
		c2 <- "result 2"
	}()
	
	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}