package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	go string은 읽기 전용 slice 이다. 언어와 표준적인 라이브러리는 string을 특수하게 처리한다. - text로 encoded 된 utf8의 컨테이너

	무슨 놈의 타입 이름이 rune이냐.. 

	아무튼 이걸 왜 쓰냐.. utf-8을 쉽게 제어하기 위해 사용한다.
	
	뭔 말인지 모르겠다.. 일단 넘어가겠음..

*/


func main() {
	const s = "สวัสดี"

	fmt.Println("Lens :", len(s))

	for i :=0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("\nRune count:", utf8.RuneCountInString(s))
}