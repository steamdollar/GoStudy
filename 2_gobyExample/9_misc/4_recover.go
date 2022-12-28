package main
import "fmt"

/*
	panic에서 recover 함수를 통해 recover하는것도 가능하다.
	recover는 panic이 프로그램을 중단하는걸 막고, 실행을 속행할 수 있게 해준다.
	
	예를 들어 서버가 에러 하나로 셧다운 된다던가 하는 상황을 막을 때..
	서버를 내리지 말고 커넥션을 끊는 식으로
*/

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	
	/*
		recover는 deferred 함수에서 실행되어야 한다.
		이를 포함한 함수가 패닉 > defer가 활성화되어 recover를 실행할 수 있도록..
	*/
	
	mayPanic()
	
	fmt.Println("After mayPanic()")
	// 이 코드는 실행되지 않음
	// 윗줄에서 panic이 나와 main()의 실행이 멈추고, defer로 넘어간다.
}

