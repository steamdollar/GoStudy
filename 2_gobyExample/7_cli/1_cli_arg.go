package main

import (
	"fmt"
	"os"
)

/*
	프로그램 실행을 parameterize할 때는 cli arg를 사용한다.
	
	e.g. go run hello.go는 run, hello.go arg를 go program에 대해 사용한다.
*/

func main() {		 
	//os.Args는 raw cli arg에의 접근을 제공해준다.
	argsWithProg := os.Args
	// 1st param은 program까지의 path, os.Arg[1:]는 프로그램에의 변수를 holding
	argsWithOutProg := os.Args[1:]
	
	arg := os.Args[3]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithOutProg)
	fmt.Println(arg)
}

// go build 후, build된 파일과 매개변수를 아무거나 넣어서 실행해보면 된다.

/*
	$ go build 1_cli_arg.go
	$ ./1_cli_arg a b c d
/*