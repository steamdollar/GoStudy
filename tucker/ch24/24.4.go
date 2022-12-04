package main
import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	
	for i := 0; i< 10; i++ {
		go func() {
			for {
				DepositAndWithDraw(account)
			}
			wg.Done()
		}()

	}
	wg.Wait()
}

func DepositAndWithDraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("negative value : %d"))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

// 이론적으로 이 함수는 무한히 1000을 넣었다 뻇다를 반복해야함.

// 근데 실행하면 조금 이따가 에러를 발생시키며 종료된다.

/*

mutex, lock을 사용해 하나만 접근 가능하도록 하는데

문제점이 있다.

1. 동시성 프로그래밍으로 인한 성능 향상을 얻을 수 없다.

심지어 과도한 락킹으로 성능이 하락되기도 한다.

다른 고루틴은 대기하고 있고 하나만이 메모리에 접근해 코드를 실행할 수 있으므로..

(Lock의 획득, 반납 자체가 성능을 사용하므로..)

2. 고루틴을 완전히 멈추게 만드는 데드락 발생 (deadlock)

고루틴이 두 개 이상의 메모리에 접근해야 할 때, 각각 다른 고루틴이 메모리에 접근해

락을 건다면 더 이상 진행이 되지 않는다.

24.5 참조

*/