package main
import (
	"fmt"
	"sync"
)

/*
	atmoic counter의 경우보다 조금 더 복잡한 state를 컨트롤하려면 mutex를 사용할 수 있다.
	mutex를 이용해 여러 개 고루틴을 이동하뎌 데이터를 안전하게 엑세스할 수 있다.
*/

type Container struct {
	mu sync.Mutex
	counters map[string]int
}
// container는 counter의 map, mutex를 포함한다.
// mutex는 복사되어선 안되므로 struct를 여기저기서 사용하려면 포인터를 사용한다.

func (c *Container) inc(name string) {
	c.mu.Lock()
	// counter에 접근하기 전 mutex를 lock
	// 다른 고루틴은 이 데이터에 접근할 수 없음
	defer c.mu.Unlock()
	// .함수가 마무리 된 후 unlock
	c.counters[name]++
}

func main() {
	c := Container {
		counters: map[string]int {"a":0, "b":0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	// 여러 개 고루틴을 동시에 돌린다.
	// 동일한 Conatainer에 접근하며, 이 중 둘은 동일나 counter에 접근한다.

	wg.Wait()
	fmt.Println(c.counters)
}