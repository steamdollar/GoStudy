package main

import "fmt"

/*
	go는 익명 함수를 지원해 closure를 형성할 수 있다.
*/

func intSeq() func() int {
	// intSeq는 다른 함수를 리턴하는데,
	// 이 리턴되는 함수는 intSeq 의 body 안에서 익명으로 정의된다.
	// 그 함수가 본 함수에 있는 데이터를 가져다 사용하면 해당 익명 함수는 closure로 바뀌게 된다.
	
	i := 0
	return func() int {
		i++
		return i
		// 
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 함수가 종료되어도 함수 안에 있는 변수가 사라지지 않고 계속 유지된다.
	// 이 점을 활용하면 전역 변수를 불필요하게 만들지 않으면서도 변수를 공유할 수 있게 된다.
	// 함수를 여러 번 호출했을 때의 결과를 쉽게 얻을 수 있다.

	newInts := intSeq()
	fmt.Println(newInts())
}