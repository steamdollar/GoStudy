package main
import (
	"fmt"
	"os"
	"strings"
)

var fl = fmt.Println

// unix program에 config 정보를 전달하는 메커니즘이 환경 변수



func main() {
	
	// k-v pair를 설정하려면 os.Setenv를 사용한다.
	os.Setenv("FOO", "1")
	
	// 환경 변수 값을 가져오려면 Getenv를 사용한다.
	fl("FOO:", os.Getenv("FOO"))
	fl("BAR:", os.Getenv("BAR"))
	
	fl()
	
	// os.Environ을 이용해 모든 k-v pair를 불러올 수 있다.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fl(pair[0])
	}
}

/*
	환경 변수를 설정하려면 실행 전에 넣어줘야 함
	
	e.g. BAR=2 go run ~~
*/