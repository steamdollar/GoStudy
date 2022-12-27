package main
import "fmt"

/*
	기본적으로 채널에서 send와 receive는 blocking이다.
	하지만 select와 default를 사용해 nonblocking으로 이를 처리할 수 있다.
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	default :
		fmt.Println("no msg received")
	}
	// non blocking receive의 예시
	// messages가 available 하다면 select는 위쪽 case를 선택, vice versa
	
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages", msg)
	default:
		fmt.Println("no msg sent")
	}
	// non blocking send의 예시, 이 쪽도 비슷하게 작동한다.
	// 여기서 msg는 messages 채널로 전송될 없다.
	// 해당 채널이 receiver없는 nonbuffer 채널이기 때문.
	// 따라서 디폴트 케이스가 선택된다.

	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	case sig := <- signals :
		fmt.Println("received signals", sig)
	default:
		fmt.Println("no activity")
	}
	// case 여러 개를 쓸 수 있음..
	// multi way non blocking select..
}