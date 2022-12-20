/*

receiver argument를 갖는 함수라고 생각하면 된다.

Go는 객체 지향적 언어가 아니기 때문에 Class라는 것이 없다.

method는 JS 등 객체지향적 프로그램을 사용할 때 class 내부의 함수를 뽑아 쓰는 것과 비슷하다.

func (var type) methodName (para1 type) {} 과 같은 형식으로 쓰면 된다.

*/


package main
import "fmt"

type emp struct {
	name string
	address string
	age int
}

func(e emp) display() {
	fmt.Println(e.name)
}

func main () {
	empdata := emp{"lsj", "seoul", 29}

	empdata.display()
}