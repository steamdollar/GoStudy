package main
import "os"

/*
	panic : 뭔가 에러가 났을 때..
	일반적인 구동시 나오면 안되는게 나왔을 경우 빠르게 fail 시키는 것이 목적이다.
*/

func main() {
	panic("a problem")
	// 
	
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	// 함수가 우리가 예상치 못한 에러를 리턴한다면 중단시키는게 panic의 주된 용도이다.
	// 예시는 새로운 파일을 생성할 때 예상못한 에러를 만나면 panic하는 것
}

/*
	이 프로그램을 실행하면 panic을 유발하고, 에러 메시지와 고루틴을 추적한 후, 0 아닌 status로 exit한다.
	panic을 만나면 이후의 코드에 도달하지 못하므로 아래 panic을 보고 싶다면 위 쪽을 제거해야함.
	
*/