package main

import ("fmt")

func main() {

	// console.log
	fmt.Println("Hello Golang\n")

	// 변수 선언

	var qwe int
	qwe = 10
	fmt.Println(qwe)

	var asd string
	asd = "go to study go"
	fmt.Println(asd)

	var wer, sdf = 123, 456
	fmt.Println(wer)
	fmt.Println(sdf)

	// 선언, 대입을 한 번에 하고 싶다면 := 사용

	xcv := "TruE"
	fmt.Println(xcv)

	// Constants
	const q = 20
	fmt.Println(q)

	// for loop - Go엔 for문만 존재한다. while, do while do not exist

	for i := 1; i <=5; i++ {
		// 조건들에 괄호 안 씀
		fmt.Println(i)
	}


	// 조건문
	con := 40
	if con < 30 {
		fmt.Println ("con < 30", ",","con :", con)
	} else {
		fmt.Println("failed")
	}
	// else문 줄 바꾸면 안됨. if문 닫는 중괄호 바로 옆에 써야 함.

	con2 := 50

	if con < 10 {
		fmt.Println ("con < 10")
	} else if con2 >= 10 && con2 <= 50{
		fmt.Println("10 < con < 50 ")
	} else {
		fmt.Println("con > 50")
	}
	

	// switch
	a, b := 2, 3

	switch a+b {
	case 1 : 
		fmt.Println("Sum 1")
	case 2 : 
		fmt.Println("Sum 2")
	case 3 :
		fmt.Println("Sum 3")
	default :
		fmt.Println("default?")
	}


	// Arrays - fixed size, elements of same type
	array1 := [3] string {"a", "b", "c"}

	fmt.Println(array1)
	fmt.Println(len(array1))

	// 이렇게 해도 되고

	var array2 [3] int
	array2[0] = 1
	array2[1] = 2
	array2[2] = 3
	fmt.Println(array2[1])
	fmt.Println(array2)


	// slice & append function

	// slice는 배열의 부분 집합 반환

	subArray2 := array2[0:2]
	fmt.Println(subArray2)

	// 이런 부분 집합을 땡겨와서 요소를 바꾸면 원본 배열 요소가 바뀜 ㄷㄷ

	subArray2[0] = 99

	fmt.Println(subArray2)
	fmt.Println(array2) // [99, 2, 3] 부왁 이게 말이 되노? ㄷㄷ

	// append

	arrayInt := [5] string {"ba", "bs", "bd", "bf", "bg"}
	sliceInt := arrayInt[1:4]

	arrayString := [5] string {"a", "as", "ad", "af", "ag"}
	sliceString := arrayString[1:3]

	fmt.Println(sliceInt)
	fmt.Println(sliceString)

	sliceInt = append(sliceInt, "text")
	fmt.Println(sliceInt)

	sliceInt = append(sliceInt, sliceString...)
	fmt.Println(sliceInt)

	// ... 을 뒤에 붙여야 문제가 없다.
	// 이 ...이 뭐지? 깊은 복사인가?
	// 파라미터로 입력된 타입의 데이터를 몇 개 받을지 확실치 않을 때 사용..

	// 함수

	para1, para2 := 10, 15

	sum, diff := calc(para1, para2)
	fmt.Println(sum, diff)

}

func calc (num1 int, num2 int) (int, int) {
	sum := num1+num2
	diff := num1-num2
	return sum, diff
}
