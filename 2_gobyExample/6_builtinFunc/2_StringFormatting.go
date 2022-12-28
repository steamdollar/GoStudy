package main
import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var f = fmt.Printf

func main() {
	p := point{1,2}
	f("struct1 : %v\n", p)
	// v는 verb를 의미 general한 Go value를 포맷팅..
	
	f("struct2 : %+v\n", p)
	// key를 함께 출력하고 싶다면 %+v
	
	f("struct3: %#v\n", p)
	// %#v는 value의 syntax까지 출력해줌
	
	f("type: %T\n", p)
	// T는 type을 의미
	
	f("bool : %t\n", true)
	// formatting boolean
	
	f("int : %d\n", 123)
	// int formattin엔 다양한 방법이 있지만
	// 표준적으로 %d를 사용, 10진수로 출력
	
	f("bin : %b\n", 14)
	// 정수의 2진수 출력
	
	f("char : %c\n", 33)
	// 정수에 대응하는 문자를 출력 (정수를 문자화하는게 아님)
	
	f("hex : %x\n", 456)
	// 정수를 16진수로 출력
	
	f("float1: %f\n", 78.9)
	// float의 소수점 자릿수를 맞춰줌
	
	f("float2 : %e\n", 123400000.0)
	f("float3 : %E\n", 123400000.0)
	// 큰 수를 한자릿수 x 10의 몇 제곱 형식으로 표현
	
	f("str1 : %s\n", "\"string\"")
	f("str2 : %q\n", "\"string\"")
	// s : basic string print , q : double-quote string
	f("str3 : %x\n", "hex this")
	// string을 16진수로 나타낼 수도 있다.
	
	f("pointer :%p\n", &p)
	// data의 pointer 출력
	
	f("width1: a%6db%6dc\n", 12, 345)
	// 글자간 간격을 주고 싶을 때 %nd사용
	// 양수는 오른쪽, 음수는 왼쪽 정렬
	
	f("width2: a%6.2fb%6.2fc\n", 1.2, 3.45)
	// float의 간격을 조절하는 것도 가능하다. 자릿수도..
	
	f("width4 : a%6sb%6sc\n", "qwe", "asd")
	
	s := fmt.Sprintf("sprintf : a %s", "string")
	fmt.Println(s)
	// Sprintf는 출력하지 않고 문자열을 반환만 한다.
	// 그 밑에서 println으로 출력함
	
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	// Fprintf는 format+print
}