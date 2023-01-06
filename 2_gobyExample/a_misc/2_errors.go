package main
import (
	"fmt"
	"errors"
)

/*
	go에서는 에러를 명시적이고, 분리된 리턴값으로 표시해준다.
	이러면 어떤 함수가 에러를 리턴하는지 쉽게 확인할 수 있다.
*/

func f1(arg int) (int, error) {
	// 관례적으로 에러는 마지막 리턴값이며 error type을 가진다.
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
		// errors.New 함수가 에러 메시지를 생성해 전달
	}

	return arg+3, nil
	// error 자리에 nil이 있다면 에러가 없다는 것을 의미
}

type argError struct {
	arg int
	prob string
}
// Error() method를 implement해 커스텀 에러 타입을 만드는 것도 가능하다.

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work w/ it"}
	}
	return arg +3, nil
}
// 이 케이스에서는 새로운 struct를 만들기 위해 &argError syntax를 사용했다.

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			// f1(i)의 리턴값인 r,e에 대해 e가 nil이 아니라면..
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}

	}
	for _, i := range []int{7,10,42} {
		if r, e := f2(i); e !=nil {
			fmt.Println("f2 failed : ", e)
		} else {
			fmt.Println("f2 worked : ", r)
		}
	}
	// 각 반복문은 error return 함수를 테스트한다.
	// 이런식으로 if line에서 line error check를 사용하는 방식을 go에서 자주 사용한다.

	_, e := f2(42)
	fmt.Println("e : ",e.(*argError))
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	// 커스텀 에러 안의 데이터를 사용하고 싶다면
	//  type assertion을 통해 인스턴스 형태로 가져와야 한다.
}