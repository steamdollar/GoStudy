package main
import "fmt"

func sample() {
	fmt.Println("Inside the Sample()")
}

func main() {
	defer sample()
	// defer : 붙은 함수의 호출을 연기 > defer를 포함한 함수가 다 실행될때까지..
	// 여기선 main이 다 실행된 후, defer가 붙은 sample()이 실행됨
	// 여러 개의 defer가 있다면 LIFO
	fmt.Println("Inside the main()")
	
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	// 3, 2, 1, Inside Sample 순서..
}