package main

import ( "fmt" )

/*
	range는 다양한 군집형 데이터의 요소들을 
	반복적으로 이용하고 싶을 때 사용한다.
*/

func main() {
	// 1. slice, array
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		// fmt.Println(a)
		// range를 이용해 어떤 반복적인 작업을 수행할 때 넣어줄 2가지의 인수가 있는데,
		// 순서대로 index와 value 이며, index를 생략하고 싶다면 _를 넣어주면 된다.
		sum += num
	}
	fmt.Println(sum)
	// 9 

	for i , num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		} else {
			fmt.Println("depreciated")
		}
	}
	// depreciated index1 depreciated
	
	// 2. map
	obj := map[string]string{"a" : "apple", "b" : "bus"}

	for k, v := range obj {
		// key, value를 이용해 작업을 수행한다.
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range obj {
		fmt.Println("key :", k)
	}

	obj2 := map[int]string{1:"a", 2:"b"}

	for k2, v2 := range obj2 {
		fmt.Printf("%d -> %s\n", k2, v2)
	}

	for i, c := range "go" {
		fmt.Println(i,c)
	}
	// string 자체도 반복이 가능하다. 일단은 이렇게만 알아두자.
}