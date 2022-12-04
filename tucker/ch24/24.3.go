package main
import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

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

이런게 동시성 문제..

고루틴이 10개를 만들었는데, 얘들이 하나의 데이터에 접근을 한다.

account.Balance += 1000

account.Balance = account.Balance + 1000

대입 연산자와 + 총 2개의 연산자가 있다.

루틴 1,2 는 add를, 3은 mov를 실행하고 있다고 하자.

..

하나의 메모리 자원에 여러 고루틴이 접근하면 무조건 문제가 생긴다.

이를 방지하는 방법이 있다.

메모리 자원을 하나의 고루틴에서만 접근하게 해야 한다.

가장 간단한 방법은 Lock을 이용하는 것인데,

이를 뮤텍스 (mutual exclusion : Mutex) 라고 한다.

한 고루틴이 락을 걸면 언락 전까지 다른 고 루틴이 접근 할 수 없음.

 24.4 에서 사용햅좌.
*/