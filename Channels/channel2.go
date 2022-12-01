/*

sender가 receiver에게 더 이상 전송할 데이터가 없다는 걸 channel을 close함으로써 알려줄 수 있다.

보통 반복문이 있을 때 이를 사용한다.

close(chan_name)

채널이 닫혔는지 확인하기 위해 다음과 같이 변수(boolean)를 추가해줄 수도 있다.

var_name, status := <- chan_name

goroutine 들간의 상호작용을 위해 channel을 사용할 수도 있다. 

한 goroutine에서는 데이터를 channel로 push, 다른 하나는 그 데이터를 받으면 된다.

*/

package main
import "fmt"
import "time"

func add_to_chan (ch chan int) {
	fmt.Println("Send data")
	for i := 0; i < 10; i++ {
		ch <- i
		// push data to channel
	}
	close(ch)
}

func fetch_from_chan (ch chan int) {
	fmt.Println("Read data")
	for {
		x, status := <- ch

		if status == true {
			// status가 false면 채널이 종료된 것
			fmt.Println(x)
		} else {
			fmt.Println("channel closed")
			break;
		}
	}
}

func main () {
	ch := make(chan int)
	go add_to_chan(ch)
	go fetch_from_chan(ch)

	time.Sleep(2 * time.Second)
	fmt.Println("Inside main()")
}

/*

1. int를 전달하는 채널 변수 선언

2. 서브루틴에서 채널을 매개변수로 갖는 add_to_chan 함수실행

3. add_to_chan에서는 채널을 통해 i(int)를 전달하고, 다 전달하면 channel close

4. fetch_from_chan 에서는 채널을 통해 들어오는 데이터의 값과 status를 받음

데이터가 전달될 때 마다 status를 읽어 status가 open이면 콘솔에 출력,

close라면 해당하는 메시지를 출력후 break

5. 메인함수는 이 둘을 실행할 시간을 2초 주고 종료

*/