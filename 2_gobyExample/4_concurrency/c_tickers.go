package main
import (
	"fmt"
	"time"
)

/*
	timer가 일정 시간 후 한 번 뭔가를 실행하고 싶은 것이라면
	ticker는 일정 간격으로 뭔가를 반복 실행하고 싶을 때 사용한다.	
*/

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	
	go func () {
		for {
			select {
			case <- done:
				return
			case t := <- ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// timer와 마찬가지로 ticker도 중지시킬 수 있다.
	done <- true
	fmt.Println("ticker stopped")
}