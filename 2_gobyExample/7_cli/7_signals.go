package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
	Unix signal들을 Go 프로그램에서 컨트롤하고 싶을 때가 있다.
	서버를 끄고 싶다던가..
*/

func main() {
	// go signal notification은 os.Signal value를 채널을 통해 전송함으로써 작동한다.
	// notification을 수신할 채널을 만든다.
	sigs := make(chan os.Signal, 1)
	
	// signal.Notify는 채널이 특정 signal의 notification을 수신하도록 등록한다.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	// main 함수에서 sigs로부터 수신할 수 있지만 
	// 좀더 실전적 환경을 가깝게 분리된 고루틴에서 어떻게 수행되는지 확인하자.
	done := make(chan bool, 1)
	
	go func() {
		// 이 고루틴은 signal을 받을 때까지 blocking 수행
		// signal을 받으면 출력하고, program이 종료를 notify
		
		sig := <- sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
	컨트롤C가 SIGINT signal이다.
	실행 후, 이걸 눌러보면 됨.
*/