package main
import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	// *int는 int pointer를 의미한다.
	
	*iptr = 0
	// pointer가 가리키는 값을 조작한다.
}

func main() {
	i := 1

	fmt.Println("initial : ", i) // 1
	fmt.Println("pointer : ", &i) // 0x~~

	zeroval(i)
	fmt.Println("zeroval : ", i) // 1
	// 값이 변하지 않는다.
	// 왜냐하면 처음 선언한 i에 대한 정보가 있는 메모리를 조작하는 것이 아니라,
	// 그 값을 복사해 다른 메모리에 저장 후, 그 값을 조작하기 때문..

	zeroptr(&i)
	fmt.Println("zeroptr : ", i) // 0
	// 애는 바뀐다. 메모리 주소를 직접 찾아가서 거기에 있는 값을 조작했기 때문
	
	fmt.Println("pointer : ", &i) // 0x~~
}