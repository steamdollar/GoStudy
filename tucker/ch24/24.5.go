
package main
import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "fork", "spoon")
	go diningProblem("B", spoon, fork, "spoon", "fork")

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstname string, secondname string) {
	for i :=0; i< 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합ㄴ다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstname)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondname)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

/*

deadlock 발생.. 모든 고루틴이 sleep 상태에 들어갔다.

차라리 이렇게라도 띄워주면 다행이지,

고루틴 여러 개중 일부만 이런 식으로 당착 상태에 빠져 sleep 상태에서 벗어나지 못한다면

이땐 deadlock 메시지를 띄워주지도 않아 에러를 찾는것이 어려울 것이다.

여기선 fork, spoon 순서만 맞춰주면 데드락을 피할 수 있지만..

작은 범위내에서, 확실할 때만 사용하는 거시 좋다.



*/