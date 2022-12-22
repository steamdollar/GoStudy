package main

import "fmt"

/*
	1.8버전부터 generic을 지원하기 시작했다.
	다른 말로는 type parameter라고도 부름

	개어려워보임 ㄷㄷ
*/

func MapKeys[K comparable, V any](m map[K]V) []K {
	// 함수명 [ 파라미터-제한조건... ] (매개변수 타입) 리턴값타입 {body...} - 이 format이 맞나?

	// MapKeys 함수는 어떤 type이든 받아서 그 keys의 slice를 리턴한다.
	// 파라미터는 K, V 두 가지가 있다.
	// K는 comparable constraint가 있는데,
	// 이건 K의 value를 == 혹은 !=로 비교할 수 있다는 것을 의미한다.
	// 이것은 Go의 map keys에서 필요한 것이다. (?)

	// v 붙은 any constraint는 특별한 제한이 없다는 것을 의미한다.

	r := make([]K, 0, len(m))
	// slice 선언, 할당 - type, length, capacity
	for k := range m {
		// 매개변수로 받은 m에 대해..
		// 생각해보면 이 range라는건 forEach에 가까운 것 같음..
		r = append(r,k)

	}
	
	return r
}

type List[T any]struct {
	head, tail *element[T]
}
// generic type의 예시, List는 타입무관 value들의 list이다.

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push (v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val : v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val : v}
		lst.tail = lst.tail.next
	}
}
// generic type도 일반적인 타입처럼 method를 정의할 수 있지만, 타입 파라미터를 잘 넣어줘야한다.
// 여기서 type은 List가 아닌 List[T]로 넣어줘야 한다.

func (lst *List[T]) GetAll() []T {
	var elems []T
	// elems의 type T를 원소로 갖는 array이다.
	for e :=lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	// e가 nil일때까지 e.val을 계속 더한다.
	return elems
	// 다 더함녀 return
}
// 함수명 앞에 붙는건 뭐지.. 제한 조건인가 이것도..
// getAll 함수는 T type의 배열을 리턴


func main () {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	// generic 함수는 invoke할 때 종종 type inference에 의존한다.
	// 즉, MapKeys를 호출할 때 K, V의 타입을 특정할 필요가 없다.
	// 컴파일러가 알아서 맞춰줄 것이다.
	// 물론 직접 명시해줘도 문제는 없다.

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}