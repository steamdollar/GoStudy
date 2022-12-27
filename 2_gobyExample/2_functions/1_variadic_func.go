package main
import "fmt"

/*
	variafic func(가변 함수)는 매개 변수의 갯수를 유동적으로 바꿀 수 있는 함수를 의미한다.
	fmt.Println도 이에 속한다.
*/

func sum(nums ...int) {
	// 매개변수의 type은 int이지만 그 갯수가 정해지지 않았을 때 ...을 붙여 이를 표현해준다.
	// 이 경우 해당 함수의 매개변수는 int의 배열과 동일하며, int의 배열을 넣어줘도 함수가 동작한다.
	fmt.Print(nums, " > ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)
}