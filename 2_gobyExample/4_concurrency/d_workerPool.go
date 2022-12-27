package main

import (
	"fmt"
	"time"
)

/*
	고루틴과 채널을 이용해 worker pool을 사용하는 방법에 관해 알아보자.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
// 몇 개의 concurrent 인스턴스를 구동할 worker가 여기 있다.
// worker들은 jobs 채널에서 일을 받아 results에 그 결과를 전송한다.
// 하나의 일이 끝나면 1초마다 sleep

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 이 worker pool을 사용하기 위해 여기서는 일감을 보내고 결과를 받아줘야 한다.

	for w := 1; w <=3; w++ {
		go worker(w, jobs, results)
	}
	// worker는 3명으로 시작, 아직 jobs이 없으므로 일단 block된 채로 시작한다.

	for j :=1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	// 여기서 job 5개를 보내고 채널을 닫는다.

	for a :=1; a <= numJobs; a++ {
		<- results
		// fmt.Println(results)
	}
	//  여기서 작업 결과물을 받는다.
	// 또한 여기서 worker 고루틴이 끝나는 것을 확정한다.
}