package main

import (
	//"bytes"
	"fmt"
	"regexp"
)

/*
	정규식에 대해 알아보자
*/

var f = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// string에 해당 패턴이 있는지 확인 - true
	// p, ch 사이에 알파벳인지?
	
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))	
	// string 패턴을 바로 사용할 수도 있지만 컴파일을 해 최적화한 struct를
	// 사용할 때도 있다.
	// true
	
	fmt.Println(r.FindString("peach punch"))
	// 이건 지정한 패턴을 만족하는 문자열을 찾아줌
	// peach
	
	fmt.Println("idx : ", r.FindStringIndex("peach punch"))
	// 만족하는 패턴 stirng을 찾아 양 끝의 index 출력 [0 5]
	
	f(r.FindStringSubmatch("peach punch"))
	// 전체 패턴에 매치되는 string, 일부만 매치되는 string을 쭉 알려줌
	// p([a-z]+ch), ([a-z]+) 를 만족하는 모든 string을 리턴
	// peach ea
	
	f(r.FindStringSubmatchIndex("peach punch"))
	 
	// 처음 하나만 찾아준다 여기까지는..
	
	f(r.FindAllString("peach punch pinch", 2))
	// 만족하는 문자열 앞에서부터 n개 찾아줌
	// -1이면 전부 찾아줌
	
	f("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// 만족하는 문자열, 서브문자열 전부 찾아줌
	
	fmt.Println(r.Match([]byte("peach")))
	// []byte 매개변수를 주고 string을 함수명에서 제거할 수도 있음
	
	r = regexp.MustCompile("p([a-z]+)ch")
	f("regexp :", r)
	// 전역 변수를 정규식을 포함해 생성할때는 MustCompile 함수를 사용한다.
	// 얘는 에러를 리턴하게 도리 상황이 오면 그냥 panic을 해버려
	// 더 안전하게 전역변수를 사용할 수 있게 해줌
	
	f(r.ReplaceAllString("a peach", "<fruit>"))
	// subset을 다른 문자열로 치환
	
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// ReplaceAllFunc 함수에 변수와 함수를 넣으면 함수를 실행해
	// 값을 바꿔줌
	f(string(out))
}