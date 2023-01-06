package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	context는 서버에서 데드라인, 신호 취소, 고루틴 등의 역할을 수행한다.
*/

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server : hello handler start")
	defer fmt.Println("server : hello server ended")
	// context.Context는 net/http machinery에 의해 각 req마다 생성되며
	// Context() method로 사용할 수 있다.
	
	select {
	// client에게 응답을 주기 전 몇 초를 기다리며,
	// 필요한 작업을 끝낸다.
	case <- time.After(5 * time.Second) :
		fmt.Fprintf(w, "hello\n")
	// 작업 진행 중 이 쪽에서는 Done signal이 오면 포착해서 작업을 취소할 수 있도록 함 
	case <- ctx.Done() :
		
		// ctx의 Err() method는 왜 Done() channel이 닫혔는지 이유를 설명해주는 error를 리턴
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	
	fmt.Println("server run 8090")
	http.ListenAndServe(":8090", nil)
}
/*
	터미널 2개 열어서
	
	하나로 서버 실행
	
	하나로 curl 요청 보냈다 취소해보면 된다.
*/