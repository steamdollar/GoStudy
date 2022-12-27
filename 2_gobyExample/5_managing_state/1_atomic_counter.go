package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Go에서 상태를 매니징하는건 여러 채널간의 의사소통이었다.
	worker pool에서 그 예시를 확인할 수 있다.

	다만 state를 매니징하는데 다른 몇 가지 옵션도 있는데 이를 알아보자.
	첫번째로 여러 개의 고루틴이 접근하는 atomic counter에 사용하ㅓ는 sync/atomic 패키지에 대해 알아보자.
*/

func main () {
	var ops uint64
	// counter

	var wg sync.WaitGroup
	// 모든 고루틴이 일을 끝낼때까지 기다리게끔 WG 가용

	for i := 0; i < 50; i++ {
		wg.Add(1)
		//counter를 천 번 증가시키는 50개의 고루틴 실행

		go func() {
			for c :=0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				// counter 값을 atomic하게 변경하는 AddUint64 사용
				// ops counter의 메모리 주소를 준다.
				// ops++ 를 사용하면 다른 값이 나옴. 모든 고루틴이 서로와 간섭하게 될테니..
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops :", ops)
	// 더 이상 ops에 쓰는 고루틴이 없다면 이제 해당 변수에 접근해도 괜찮다.
	// LoadUint64 메소드를 이용해 atmoic하게 변수를 읽는 것도 가능하다.


}