package main

import (
	"fmt"
	"os"
)

/*
	defer : 함수 호출이 프로그램의 실행 후에 (보통은 clean up을 위해) 
	실행되는 것을 보증하기 위해 사용한다.
	
	파일 생성 > 쓰기 > close 과정을 거칠 때 defer가 다음과 같이 사용될 수 있다.
*/

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	// closeFile(f)는 메인 함수 실행 종료 후 = writefile이 끝난 후 실행된다.
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating..")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f 
}

func writeFile(f *os.File) {
	fmt.Println("writing..")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	err := f.Close()
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}
}
// deferred 함수일지라도 파일을 닫을때 에러 체크를 해주는게 좋다.