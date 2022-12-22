package main
import "fmt"

/*
	channel은 고루틴들을 연결해주는 파이프라고 보면 된다.
	한 고루틴에서 채널을 통해 값을 보내고 그걸 다른 고루틴에서 받을 수 있다.
*/

func main () {
	messages := make(chan string)
	// make를 이용해 채널을 생성한다.
	// 채널은 자신이 전달하는 데이터 타입을 적어준다.

	go func() { messages <- "ping"}()
	//  다음과 같은 syntax로 채널에 데이터를 전송할 수 있다.

	msg := <- messages
	// 데이터를 받는건 이렇게 써주면 된다.
	// 화살표는 항상 <- 를 사용한다. ->는 없음
	fmt.Println(msg)
}