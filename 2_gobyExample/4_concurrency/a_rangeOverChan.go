package main
import "fmt"

/*
	기본적인 데이터를 반복하는데 for와 range를 사용했었다.
	이 신텍스를 채널에서 반복적으로 데이터를 주고 받는 것에도 사용할 수 있다.
*/

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// 채널에서 두 개의 value를 반복하고 close

	for elem := range queue {
		fmt.Println(elem)
	}
	// range는 queue로부터 받은 element 각각에 대해 반복한다.
	// 채널이 닫히면 반복도 종료된다.
}