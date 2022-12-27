package main

import (
	"fmt"
)

/*
	ref : https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81410/%EB%A7%B5-map
	map은 Go의 관계형 데이터 타입이다.
	JS의 객체처럼 key - value를 묶어 저장하고 싶을 때 사용한다.

	slice처럼 참조 타입이며, 선언과 초기화가 비슷한 방식으로 행해진다.
*/


func main() {
	var map1 map[int]string

	fmt.Println(map1)
	// map[]

	// map에 값 할당
	map1 = map[int]string {
		// key가 int, value가 string임을 의미한다.
		237:"steamdollar",
		200:"mercedes",
		112:"RevAssault",
	}

	fmt.Println(map1)

	// map의 data 확인
	// map에 궁금한 키가 존재하는지 보고 싶다면 다음과 같이 확인할 수 있다.

	fmt.Println(map1[200]) // mercedes
	fmt.Println(map1[345]) // 
	// 파라미터로 넣어준 키가 존재한다면 그 값을 알려주고, 존재하지 않는다면 공란이 출력된다.

	// 존재하는지의 여부를 true/false로도 확인하고 싶다면
	// 변수를 2개 선언해 할당해주어야 한다.

	val, isExist := map1[188]
	fmt.Println(val, isExist)

	val2, isExist2 := map1[456]
	fmt.Println(val2, isExist2)
	// 얘는 map에 해당 값이 없으므로 공란, false가 반환된다.

	_, isExist3 := map1[112]
	fmt.Println(isExist3)
	// _를 쓰면 이 값을 생략하겠다는 것을 의미한다.
	// 이 경우 true/false 값만이 출력된다.

	/* map에 data 추가 */
	// map에 데이터를 추가하고 싶다면 다음과 같이
	// map이름[key] = value
	// 형식으로 추가해줄 수 있다.

	map1[188] = "zero"
	map1[138] = "luminous"

	fmt.Println(map1)
	
	// 데이터를 삭제하고 싶다면 delete 함수를 사용한다.
	// delete(map이름, "제거할 데이터 key")
	delete(map1, 200)

	fmt.Println(map1)

	// 선언과 초기화를 한 번에 하고 싶다면 다음과 같이 해주면 된다.

	var map2 = map[string]string {
		"apple" : "red",
		"grape" : "purple",
		"banana" : "yellow",
	}
	
	fmt.Println(map2)

}