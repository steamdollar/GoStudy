package main

import (
	"flag"
	"fmt"
)

/*
	cli program에 대해 옵션을 특정할 수 있는 기본적인 방법
	
	flag package가 기본적인 cli flag parsing을 지원한다.
*/

var fl = fmt.Println

func main() {
	// 기본적인 flag 정의는 string, int, bool에 대해 가능하다.
	// 1st par : flag명, 2nd par : default, 3rd par : explain
	wordPtr := flag.String("word", "foo", "a string")
	
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")
	
	// 프로그램의 다른 부분에서 정의된 변수에 값을 대입하도록 하는 것도 가능하다.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	
	// execute cli parse
	flag.Parse()
	
	// 여기서는 flag들을 보여주기만 할것
	// 실제 option의 value들을 가져오기 위해선 dereference 해야한다.
	fl("word : ", *wordPtr)
	fl("numb : ", *numbPtr)
	fl("fork : ", *forkPtr)
	fl("svar : ", svar)
	
	// flag의 key값 지정이 따로 되지 않은 나머지 cli arg들을 출력
	// 이런 flag들을 모든 flag의 value들을 다 쓰고 마지막에 써줘야함
	fl("tail : ", flag.Args())
}