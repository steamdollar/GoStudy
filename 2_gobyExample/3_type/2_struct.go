package main
import "fmt"

type person struct {
	name string
	age int
}
// struct를 선언

func newPerson(name string) *person {
	p := person{name : name}
	p.age = 42
	fmt.Println(&p)
	return &p
}
// 함수의 scope에서 살아남으므로 지역 변수의 pointer를 리턴해도 안전하다.

func main() {

	st := person{name : "sila", age : 29}
	// type을 선언했다면 선언, 할당을 동시에 하는 것도 가능
	fmt.Println(st)
	//

	eth := newPerson("alice")
	hte := &eth
	fmt.Println(&eth)
	fmt.Println(hte)

	fmt.Println(eth)
}