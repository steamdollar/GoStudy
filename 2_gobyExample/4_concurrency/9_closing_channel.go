package main
import "fmt"

/*
	채널을 닫는다는 것은 더 이상 값이 전송되지 않는다는 뜻을 내표한다.
	채널의 recevier와 상호작용하는데 이 채널을 닫는다는 행위가 유용하다.
*/

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			// 채널이 닫히면 more이 false로 바뀐다.
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <=3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<- done
}