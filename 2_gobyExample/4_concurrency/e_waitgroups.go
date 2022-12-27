package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	여러 개의 고루틴이 끝날때까지 기다리기 위해 wait group을 사용할 수 있다
*/

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", id)
}

func main () {
	var wg sync.WaitGroup
	// 이 waitgroup은 여기서 시작한 모든 고루틴이 끝날때까지 기다린다.
	// 혹시 이 Waitgroup이 다른 함수에 명시적으로 전달된다면 포인터를 사용해야한다.

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// 몇 개의 고루틴을 렁칭하고 각가에 대해 카운터를 건다.
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
		// worker 함수를 익명함수(closure)로 감싸 worker가 끝났음을 
		// 확실히 Waitgroup에 알릴 수 있게 해주면
		// worker함수는 자신의 실행 자체에 내포된 concurrency를 신경쓰지 않아도 된다.
	}
	wg.Wait()
	// WairGroup counter가 0이 될 때까지
	// 즉, 모든 worker가 자신이 끝났음을 알릴 때까지 여기서 Block된다.
	fmt.Println("all workers done")
}