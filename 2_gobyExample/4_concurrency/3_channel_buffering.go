package main
import "fmt"

/*
	gochannel은 디폴트로 unbuffered 되어 있다.
	즉, 전송할 데이터를 받을 receiver가 준비 되어 있을 때만 전송이 가능하다.

	Buffered Channel은 한정된 갯수의 value를 대응하는 receiver 없이 전송한다.
*/

func main() {
	messages := make(chan string, 2)
	// 두 개의 값을 buffer up하는 string channel을 make

	messages <- "buffered"
	messages <- "channel"
	// buffered channel이므로 value들을 receive 없이 보낼 수 있다.

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// 이 후, value들을 평소처럼 받을 수 있다.
}