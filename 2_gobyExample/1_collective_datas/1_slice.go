/*
	go에서 array는 크기가 고정되어있고, 부분 배열을 발췌하는 기능을 가지고 있지 않는다.
	이를 보완하기 위한 것이 slice인데, 이는 array와 다르게 크기를 변경할 수 있고, 부분 발췌가 가능하다.

	다만 이런 slice는 다른 자료형과는 그 구조가 조금 다르기 때문에 선언, 초기화할 때 주의할 점이 몇 가지 있다.

	다른 변수의 경우 선언을 할 때 크기가 정해져 있으므로 (int 등은 1개, 배열은 내가 정해준 length)
	이를 저장할 메모리가 자동으로 할당이 된다.
	실제로 선언 후, 이를 println으로 출력하면 int와 string 각각 0, "" 이 나타나는 것을 확인할 수 있다.
	배열의 경우도 0, "" 로 채워진 배열이 나온다.

	그런데 slice는 다음과 같이 동일한 방식으로 선언하면 아무것도 나오지 않는다.
	이유는 slice의 크기를 모르기 때문에 컴퓨터가 얼만큼을 할당해야 하는지 알 수가 없기 때문이다.
	(length를 출력해보면 0이 나온다.)
	그렇기 때문에 slice를 선언한 후, 값도 같이 초기화를 해주어야 추가적인 조작이 가능하다.
	값이 없이 선언만 된 slice를 Nil slice라고 한다.

	슬라이스는 기본적으로 배열의 위치를 가리키는 포인터(ptr), 길이(length), 전체 용량 크기(cap)를 가지는 참조 타입이다.
	따라서 슬라이스를 복사한다는 것은 같은 주소를 참조하게끔 만들어주는 것이고, 부분 복사 또한 데이터의 일부를 참조해 오는 것이다.
	


*/

package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)
	fmt.Println(len(a))

	a = []int{1,2,3}
	fmt.Println(a)
	fmt.Println(len(a))

	a[1] = 7

	fmt.Println(a)
	fmt.Println(cap(a)) // 3
	fmt.Println(len(a)) // 3

	// 길이는 slice의 원소의 갯수이고, 용량은 이를 전부 저장하기 위해 필요한 메모리의 양이다.

	/*
		make() 함수를 이용해 슬라이스를 선언해보자.
		이를 이용해 slice를 선언하면서 크기를 미리 지정해줄 수 있다.
		바꿔말하면 값을 저장할 메모리를 미리 할당해두는 것이다.

		다음과 같이 slice type, length, capacity(생략 가능)을 파라미터도 slice를 선언할 수 있다.
	*/

	s := make([]int, 3)


	fmt.Println("emp:", s)
	fmt.Println(len(s))
	// 이 때는 빈 slice라도 length를 가지고 있다는 것을 확인할 수 있다.

	s = []int{1,2,3}
	fmt.Println(s)
	// slice에 새로운 값 할당

	/* slice append, merge, copy */

	/*
	append 함수를 이용해 slice에 data를 추가할 수 있다.
	slice 용량이 남아있을 경우엔 그 용량 내에서 슬라이스 길이를 변경할 수 있지만,
	append처럼 일반적으로 용량이 늘어나는 조작을 거칠 경우 새로운 배열을 생성하고,
	기존 배열 값들을 모두 새 배열에 복제한 후 다시 slice를 할당한다.
	*/

	for i := 4; i <= 8; i++ {
		s = append(s, i)
	}

	fmt.Println(s)

	slice1 := []int{1,2,3}
	slice2 := []int{0xa,0xb,0xc}

	slice1 = append(slice1, slice2...)
	// slice에 slice를 추가할 경우 추가할 slice 뒤에 ...을 붙여 줘야 한다.
	// 이는 slice의 모든 요소들의 set을 의미한다.

	fmt.Println(slice1)

	// copy

	/*
		copy()

		copy(데이터를 받을 slice, 복사할 slice) 포맷으로 작성하면 된다.
	*/

	slice11 := []int{0,1,2}
	slice12 := make([]int, len(slice11))

	copy(slice12, slice11)

	fmt.Println(slice12)

	/* 부분 배열 복사 */

	/* 
		slice의 일부분만 잘라서 복사를 할 수 있다. 

		slice21에 slice22의 m~n번째 index를 복사해 넣고 싶다면

		slice21 := slice22 [m:n+1] 
		
		형식으로 사용하면 된다.
	*/

	slice21 := make([]int, 0)
	slice21 = append(slice21, 1, 2, 3, 4, 5)
	fmt.Println(slice21)

	l := slice21[1:3]
	fmt.Println(l)

	m := slice21[2:6]
	m[0] = 99
	fmt.Println(m)

	fmt.Println(slice21)
	// 같은 메모리를 참조하고 있으므로 일부 배열을 발췌한 m의 요소를 수정하면
	// 같은 곳을 참조하는 원본 slice인 slice21의 값도 변한다.

	/* 2d slice */

	/*
		배열과 마찬가지로 2중 slice도 존재한다.

	*/

	slice2D := make([]([]int), 3)

	for i :=0; i < 3; i++ {
		innerLen := i + 1
		slice2D[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			slice2D[i][j] = i + j
		}
	}

	fmt.Println(slice2D)
	
}