// & operator는 변수의 address를 구하는데 사용된다.
// 즉 &a는 변수 a의 memory address를 출력해준다.

package main
import "fmt"

func main() {
	a := 20;
	fmt.Println("Address : ", &a)
	// a의 value는 메모리의 어디에 저장되어있는지를 숫자로 표현
	fmt.Println("Value:", a)

	b := &a
	fmt.Println("b :",  b)
	fmt.Println("*b :", *b)
	// b라는 a의 address에는 어떤값이 있는지를 알려줌

	*b++
	// if b = &a, then *b  = a ?
	fmt.Println(a == *b)
	fmt.Println(a)

	fmt.Println(b)
}