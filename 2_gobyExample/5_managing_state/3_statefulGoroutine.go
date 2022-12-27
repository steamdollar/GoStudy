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
	}
}