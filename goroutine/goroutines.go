/*

Go는 concurrency를 제공한다. 여러 개의 task를 동시에 실행할 수있다. 
(<> parallelism : task를 여러 개의 subtask로 쪼개 병럴적으로 실행)

이는 Goroutines, Channels를 통해 구현된다.

i) Goroutines

gorountine은 다른 함수와 동시에 실행될 수 있는 함수이다.

일반적으로 한 함수가 호출되면 그 함수가 컨트롤을 갖고, 실행이 완료되면 호출한 함수로 컨트롤이 돌아간다.

컨트롤이 다시 돌아온 후에 호출한 함수는 나머지 코드를 실행한다.

goroutine에서는 내부에서 호출된 함수가 다 실행이 끝나길 기다리지 않고 호출된 함수가 실행되는 동안

자신의 나머지 코드를 실행한다.

main 프로그램도 본인의 코드가 다 실행되면 goroutine이 다 실행되지 않고 종료된다.

함수 호출시 go 키워드만 붙이면 Goroutine을 추가할 수 있다.

*/

package main
import "fmt"

func display () {
	for i:=0; i<5; i++ {
		fmt.Println("In display")
	}
}

func main() {
	go display()

	for i:=0; i<5; i++ {
		fmt.Println("In main")
	}
}

// go routine이 시작되기 전에 main 함수가 종료되었으므로 In display는 출력되지 않는다.