package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
// basic interface for shape

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//circle, rectangle 두 개의 type이 있다.

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2* r.width + 2*r.height
}
// implement interface to struct
// 이 때 interface에 있는 모든 method를 implement해야한다.

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
// geometry -> rectangle

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
// 변수에 interface type이 있다면
// interface 안의 method들을 호출할 수 있다.
// 

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
	// circle, rect type 모두 geometry를 implement하고 있으므로
	// 이 instance들을 measure의 매개변수로 사용할 수 있다.
}