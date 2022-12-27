package main
import "fmt"

/*
	채널을 함수의 매개변수로 사용할때는 이 채널이 값을 보내는 건지 받는건지 특정할 수 있다.
	이런 특정은 프로그램의 타입 안정성을 증가시켜준다.
*/

func ping(pings chan<- string, msg string) {
	pings <- msg
}
// ping 함수는 값을 전송하는 채널만을 매개변수로 받는다.
// 값을 받는 채널을 넣으면 컴파일과정에서 에러가 난다.

func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	// pings channel로 값을 받고
	pongs <- msg
	// 그걸 pongs 채널로 전달
}
// pong 함수는 값을 전송하는 채널(pings), 받는 채널(pongs) 각각 하나씩 매개변수로 받는다.

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}