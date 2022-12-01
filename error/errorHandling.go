package main

import (
	"fmt"
	"os"
)

func fileopen(name string) {
	f, er := os.Open(name)
	// 일반적으로 함수는 마지막 리턴값에 에러를 넣어준다.

	fmt.Println(f)
	fmt.Println(er)

	if er != nil {
		fmt.Println(er)
		fmt.Println("에러가 nil이 아니면 에러 출력")
		return
	} else {
		fmt.Println("file opened", f.Name())
	}
}

func main() {
	fileopen("invalid.txt")
}
