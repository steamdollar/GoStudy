package main
import (
	"fmt"
	"time"
	// "reflect"
)

/*
	Rate limiting은 리소스, 서비스 관리에 매우 중요한 메커니즘이다.
	Go는 이를 고루틴, 채널, ticker를 이용해 지원한다.
*/

func main() {
	/* basic rate limiting */
	// 들어오는 요청에 대한 핸들링을 제한하고 싶다고 가정
	requests := make(chan int, 5)
	for i := 1; i <=5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	// fmt.Println(reflect.TypeOf(limiter)) // <-chan time.Time
	// limiter 채널은 0.2초에 한번 값을 받는다.
	// 이 채널이 이 속도 제한 기법의 조절자가 됨

	for req := range requests {
		<- limiter
		fmt.Println("req", req, time.Now())
	}
	// 각 요청을 서비스하기 전에 limiter 채널로부터의 리시브를 막음으로써,
	// 1개의 요청을 처리하는데 0.2초로 속도를 제한할 수 있다.

	////

	burstyLimiter := make(chan time.Time, 3)
	// 전체적인 rate limit를 유지한채 일시적으로 빠르게 요청을 처리하고 싶을 수도 있다.
	// 이는 limiter 채널을 버퍼링 함으로써 구현할 수 있다.
	// burstyLimiter는 3개의 event를 가속시켜줄 수 있다.

	for i := 0; i <3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t:= range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	// 0.2초마다 burstyLimiter에 값을 추가한다. (최대 3개)

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	// 추가로 5개의 요청이 들어온다면, 처음 3개는 burstyLimiter를 이용해 빠르게 처리될 수 있다.
	close(burstyRequests)
	for req := range burstyRequests {
		<- burstyLimiter
		// burstyRequests 채널의 반복에서 burstyLimiter로부터 value 방출 (x3)
		fmt.Println("request", req, time.Now())
	}
}

