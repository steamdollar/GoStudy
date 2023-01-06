package main

import (
	"fmt"
	"os"
)

/*
	program 종료시 exit state를 지정할 수 있다.
*/

func main() {
	// Exit()을 사용할 경우, defer도 작동하지 않아 이 코드는 실행되지 않음
	defer fmt.Println("!")
	
	// C와는 다르게 Go는 exit status를 정수로 보여주지 않으니 이를 사용해야한다.
	os.Exit(3)
}