/*

mutex는 mutual exclusion을 줄여쓴 말이다.

이는 하나의 리소스(변수값 이라던가)에 여러 서브루틴이 동시에 접근하는 것을 원치 않을 때 

사용할 수 있다.

mutex는 sync package에 있으며, lock, unlock 2개의 method가 있다

*/

package main
import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
)

var count = 0
// 모든 routine 인스턴스에서 접근할 수 있는 변수를 선언

func process(n int) {
	for i:=1; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		// 0~2초 주기로 Sleep
		temp := count
		temp ++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
	}
	fmt.Println("Count after i= "+ strconv.Itoa(n) + " | Count: ", strconv.Itoa(count))
	// strconv : conversions to/from string representations og basic data types
	// Itoa : int to string
}

func main() {
	for i:= 1; i < 4; i++ {
		go process(i)
	}
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count)
}

/* result

Count after i= 1 | Count:  9
Count after i= 2 | Count:  12
Count after i= 3 | Count:  14
Final Count: 14

*/

/**

원래 의도는 한 번에 10씩 올려서 최종 count를 30으로 하고 싶었는데

전혀 다른 값이 나올 뿐더러, 출력되는 순서도 불규칙하게 나온다 (할 때마다 다름)

여기서는 3개의 goroutine이 count의 값을 증가시키려고 하고 있다.

기본적인 단계는

count값을 temp로 복사 > temp를 증가 > temp값을 다시 count로 복사

한 루틴이 3번째 과정을 실행한 직후 다른 루틴이 예전 value를 저장해버렸기 때문에 값이 이상해진다.

*/