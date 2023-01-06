package main

import (
	"fmt"
	"io"
	"os/exec"
)

var fl = fmt.Println

func main () {
	// simple cmd (takes no arg or input) > just print sth to stdout
	// exec.Command helper가 이 external process를 표현하기 위한 객체를 생성한다.
	
	// shell에 date 명령어를 입력하는 것이라고 보면 된다.
	dateCmd := exec.Command("date")
	
	// Output() method는 커맨드를 실행하고 끝날때까지 기다린 후, std output을 collect
	// error가 없다면 dateOut이 date info를 담은 byte를 hold
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	
	fl("> date")
	fl(string(dateOut))
	
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		// Command() method는 명령어 실행 도중 에러가 있다면 *exec.Error를 리턴하고
		case *exec.Error:
			fl("failed executing :", err)
		// 실행은 되었지만 0 아닌 값으로 종료되면 *exec.ExitError 리턴
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}
	
	// 데이터를 외부 프로세스에 그 쪽의stdin 형태로 넣고, 결과값을 그쪽의 stdout으로 받아와보자.
	grepCmd := exec.Command("grep", "hello")
	
	// StdinPipe, StdOutput으로 grab
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	// grepout pipe만 모았지만 grepIn도 모을 수 있음
	grepCmd.Wait()
	
	// err check는 생략
	fl("> grep hello")
	fl(string(grepBytes))
	
	// 옵션들은 한번에 묶어 써줄 수도 있고, 따로 써줄 수도 있음
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	
	fl("> ls -a -l -h")
	fl(string(lsOut))
}