package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	math/rand package가 유사난수 생성
*/

var f = fmt.Print
var fl = fmt.Println

func echoline() {
	fl("=============")
}

func main () {
	// rand.Intn는 0~100 사이의 정수를 리턴
	f(rand.Intn(100), ",")
	f(rand.Intn(100))
	fl()
	
	
	// 0~1 사이의 소수 리턴
	echoline()
	fl(rand.Float64())
	
	// 이런 식으로 곱셈 덧셈을 잘 조합해 범위 조절 가능
	f((rand.Float64()*5)+5, ", ")
	fl((rand.Float64() * 5 ) + 5)
	echoline()
	
	
	// 생성마다 난수를 바꾸려면 seed를 줘야한다.
	// 비밀번호같은데 사용할 난수를 생성하고 싶으면 crypto/rand를 사용할 것
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	
	f(r1.Intn(100), ", ")
	fl(r1.Intn(100))
	
	echoline()
	
	s2 := rand.NewSource(42)
	// NewSource 메소드를 이용해 다시 한 번 난수화 해줄 수 있다.
	r2 := rand.New(s2)
	
	f(r2.Intn(100), ", ")
	fl(r2.Intn(100))
	
	echoline()
	
	// 동일한 값을 준다면 동일한 값이 나온다. 이딴게.. 난수..?
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	
	f(r3.Intn(100), ", ")
	fl(r3.Intn(100))
	
	echoline()
	
	// NewSource에 다른 값을 주면 다른 결과물이 나온다.
	s4 := rand.NewSource(90)
	r4 := rand.New(s4)
	
	f(r4.Intn(100), ", ")
	fl(r4.Intn(100))
}