package main
import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
	mutex는 여러 개의 고루틴이 공유하는 state를 명시적으로 lock했다.
	다른 방법으로, 고루틴과 채널의 built-in syncronization 특성을 이용하는 것이 있다.

	이 채널 기반 접근은 커뮤니케이션을 이용해 메모리를 공유하고, 
	각 데이터의 조각을 정확히 1개의 고루틴에 의해 소유하도록 하는
	Go의 아이디어와도 방향성을 같이한다.

*/

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}
/*
	여기서는 state가 하나의 고루틴에 의해 소유된다.
	이는 다른 동시적 접근에 의해 데이터가 오염될 일이 없다는 것을 보장한다.
	이 state를 w,r 하기 위해 다른 고루틴이 소유 고루틴에 메시지를 보내고 응답을 받을 것이다.
	이 readOp, writeOp struct는 그 요청을 캡슐화하고, 소유 고루틴이 응답할 방식을 제공한다.
*/

func main() {
	var readOps uint64
	var writeOps uint64
	// 작동의 횟수 카운트

	reads := make(chan readOp)
	writes := make(chan writeOp)
	// 채널들은 다른 고 루틴이 w, r 요청을 보낼때 사용한다.

	go func () {
		var state = make(map[int]int)
		
		for {
			select {
			case read := <-reads:
				// reads 채널에서 요청을 받으면
					read.resp <- state[read.key]
					// 응답으로 read.resp 채널에 state의 값을 전송
			case write := <-writes:
				// write 채널에서 요청을 받으면
				state[write.key] = write.val
				// 받은 값으로 state를 바꾸고
				write.resp <- true
				// true를 응답으로 전송
			}
		}
	}()
	// state를 소유하는 고루틴
	// 응답은 요청된 동작을 수행하고, 응답 채널에 성공을 의미하는 응답을 전송하는 방식으로 실행된다.
	
	for r := 0; r <100; r++ {
		go func() {
			for {
				read := readOp{
					key : rand.Intn(5),
					resp : make(chan int)}
				// read 변수 선언 후
				reads <- read
				// 이를 채널에 넣는다.
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
				// readOps 값에 1을 더하고 1ms sleep
				
			}
		}()
	}
	// 고루틴 100개를 실행 > reads 채널을 통해 state 소유 고루틴에게 reads를 이슈
	// 각 read는 readOp을 생성하고 이를 reads 채널을 통해 전송한 다음, resp 채널을 통해 결과를 받는다.
	
	for w :=0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key : rand.Intn(5),
					val : rand.Intn(100),
					resp : make(chan bool)}
				// write 변수 선언
				writes <- write
				// 변수를 채널로 전송
				<- write.resp
				// 응답 수신
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// write하는 고루틴 10개도 동일하게 작동한다.
	time.Sleep(time.Second)
	
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)
}