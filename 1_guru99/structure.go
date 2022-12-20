/*

structure는 사용자 정의 데이터 타입 (class랑 비슷하다고 보면 됨)

*/

package main

import "fmt"

type emp struct {
	name string
	address string
	age int
}

func display(e emp) {
	fmt.Println(e.name)
}

func main() {

	empdata1 := emp{name : "lsj", address : "seoul", age : 29}
	fmt.Println(empdata1)
	display(empdata1)
}
