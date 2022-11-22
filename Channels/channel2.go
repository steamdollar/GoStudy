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