package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*
	net/http package를 이용해 http 통신을 할 수 있다.
*/

func main() {
	
	// 지정된 url에 http get req를 전송
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	fmt.Println("response status : ", resp.Status) // 200 OK
	
	scanner := bufio.NewScanner(resp.Body)
	// response의 첫 5줄만 출력
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}