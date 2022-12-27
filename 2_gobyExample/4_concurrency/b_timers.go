package main
import (
	"fmt"
	"time"
)

/*
	go 코드를 즉시가 아닌 조금 후에 실행하거나 
	일정 간격으로 반복적인 실행을 하고 싶을 때가 있다.
	go에서는 timer, ticker를 이용해 이를 구현한다.
*/

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	// timer는 미래의 single 이벤트를 나타낸다.
	// timer에 얼마나 기다릴지 알려주거나, 채널을 줘서 원하는 시간대에 알려줄 수도 있다.
	
	<-timer1.C
	// timer가 실행되었다는 value를 전송할때까지
	// 이 코드는 timer의 채널 C를 block한다.
	
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func () {
		<- timer2.C
		// 여기서 1초간 막혀있는 동안 아래 Stop()이 실행되서 fired가 출력되기 전
		// stopped가 출력된다.
		fmt.Println("Timer 2 fired")
	}()
	// 

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	// 단순히 기다리기만을 원한다면 time.Sleep을 사용해도 좋지만
	// timer는 작동하기전에 멈추는 것도 가능하다.

	time.Sleep(2 & time.Second)

}