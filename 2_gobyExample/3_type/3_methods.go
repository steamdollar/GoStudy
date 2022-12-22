package main
import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}
// area method는 *rect type receiver가 있다.

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}
// method는 pointer나 value 어떤 값이든 받을 수 있다.
// 이는 value receiver를 가진 method이다. 

func main() {
	r := rect{width: 10, height :5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	// struct를 포인터를 받는 area method에 넣었지만 문제 없이 실행됨 

	rp := &r

	fmt.Println("area : ", rp.area())
	fmt.Println("perim :", rp.perim())
	// method 호출시 go는 자동으로 value와 pointer를 맞춰준다.
	// method call이 복사되는 것을 피하거나 받은 struct를 변형하기 위해 pointer를 받을 수 있다.
}