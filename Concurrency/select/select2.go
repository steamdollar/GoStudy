package main

import "fmt"
import "time"

func data1 (ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from data1()"
}

func data2 (ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from data2()"
}

func main() {
	chan1 := make (chan string)
	chan2 := make (chan string)

	go data1(chan1)
	go data2(chan2)

	// 이 select문이 실행 될 때, chan1, chan2로부터 데이터를 아직 받지 못했지만,
	// 메인 함수는 기다리지 않는다. 디폴트 케이스가 실행된 후, select문을 빠져나간다.

	// time.Sleep(3 * time.Second)
	// 만약 딜레이를 줘서 두 채널 모두로부터 데이터가 수신된 상태라면 랜덤으로 출력
	select {
		case x := <- chan1:
			fmt.Println(x)
		case y:= <- chan2:
			fmt.Println(y)
		default:
			fmt.Println("Default case exe")
	}

}