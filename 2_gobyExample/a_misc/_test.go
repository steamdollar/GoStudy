package main
//package main이 아니라도 상관 없음


import (
	"fmt"
	"testing"
)

/*
	코드를 테스트 하고 싶다면 testing package를 사용하고
	run 대신 test를 사용해 프로그램을 실행한다.
*/

// 이 대소비교하는 함수를 테스트해보자.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test 로시작하는 함수를 작성해 테스트를 생성한다.
func TestIntMinbasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error*는 에러를 리포트해주지만 코드는 계속 실행
		// t.Fatal*는 에러를 리포트하고 코드를 바로 중단
		t.Errorf("Intmin(2,-2) = %d; want -2", ans)
	}
}

// 여러 개체에 대해 반복적인 test를 하려면 table driven style을 사용한다.
func TestIntMinTableDriven(t *testing.T) {
	var tests = [] struct {
		a,b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0 , -1}
	}
	
	for _, tt := range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		// t.Run을 이용해 각각의 table entry에 대한 subtest를 실행할 수 있다.
		// go test -v 를 입력하면 각각 따로 터미널에 나타남
		t.Run(testname, func(t *testing.T){
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want{
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark는 _test.go 파일 내부의 go를 테스트하며,
// 이름이 Benchmark로 시작한다.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i <b.N; i++ {
		IntMin(1,2)
	}
}