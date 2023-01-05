package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	line filter는 imput을 stdin으로 읽고, 처리하고 ,stdout으로 출력해주는
	일반적인 프로그램 타입이다.	
	
	$ echo 'hello'   > /tmp/lines
	$ echo 'filter' >> /tmp/lines
*/

func main () {
	// unbuffered os.Stdin을 buffered scanner로 wrapping해 scan method를 사용할 수 있다.
	scanner := bufio.NewScanner(os.Stdin)
	
	// Text는 current token을 return, 처리(이 경우 대문자로 변환), 다음줄로 이동, 반복
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	
	// scan 도중 에러를 체크.
	// EOF는 scan에 의해서는 error로 간주되지 않음
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	
}