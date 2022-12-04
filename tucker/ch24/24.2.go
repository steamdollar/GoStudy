package main

import (
"sync"
"fmt"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i<=b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d 까지 합계는 %d 입니다\n",a ,b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++{
		go SumAtoB(1, 1000000)
	}

	wg.Wait()
	// 얘가 있을 때 없을때를 비교
}