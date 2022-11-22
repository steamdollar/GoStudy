/*

select

select는 channels에서 작동하는 switch하고 볼 수 있다. 

channel의 동작이 case에 대응한다.

특정 채널이 읽히면 그에 연결된 case가 실행되고

여러 개의 case가 준비되면 무작위로 하나가 실행된다.

어떤 case도 준비되지 않았을때의 경우 디폴트 case를 설정 할 수도 있다.

*/

package main
import "fmt"
import "time"

func data1(ch chan string) {
	time.Sleep( 2 * time.Second)
	ch <- "from data1()"
}

func data2 (ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from data2()"
}

func main() {
	// create channel variables
	chan1 := make(chan string)
	chan2 := make(chan string)

	// invoke subroutines with channel variable
	go data2(chan2)
	go data1(chan1)

	// select 문에서 chan1, chan2로 부터의 데이터를 기다림
	// 먼저 오는 쪽의 case를 실행 후, block에서 exit
	select {
	case x := <- chan1 :
		fmt.Println(x)
	
	case y := <- chan2 :
		fmt.Println(y)
	}
}