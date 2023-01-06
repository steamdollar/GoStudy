package main

import (
	"os"
	"os/exec"
	"syscall"
)

/*
	현재 Go press를 다른 걸로 대체하고 싶을 때가 있다. (go 아닌 다른 파일로도)
	이럴 땐 exec 함수를 사용한다.
*/

func main() {
	// ls를 exec 할건데, go는 우리 실행하길 원하는 binary의 절대경로를 요구하므로
	// exec.lookPath method를 이용해 해당 명령어를 찾는다.
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	
	// Exec는 arg들을 slice로 받는다.
	args := []string{"ls", "-a", "-l", "-h"}
	
	// Exec를 사용하기 위해 환경 변수들도 필요하다.
	// 여기선 그냥 현재 환경 변수를 사용하면 된다.
	env := os.Environ()
	
	// 여기가 실제로 syscall.Exec이 호출되는 곳이다.
	// 이 호출이 성공적이라면 우리 process는 여기서 끝나고 지정한 process로 대체된다.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}