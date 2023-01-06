package main

import (
	"fmt"
	"net/http"
)

// net/http 서버를 이용해 간단한 http server를 구축할 수 있다.

// net/http 의 server는 handler인데,
// handler는 http.Handler interface를 구현하는 객체이다.
func hello(w http.ResponseWriter, req *http.Request) {

	// response writer가 http로 돌려줄 데이터
	fmt.Fprintf(w, "hello\n")
}

// 모든 http req header를 읽어 response body에 echo
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	
	fmt.Println("server run 8090")
	http.ListenAndServe(":8090", nil)
}