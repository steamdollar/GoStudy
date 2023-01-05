package main

import (
	"bufio"
	"fmt"
	"os"
)

func check( e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dump string into file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	
	// 세세하게 단계를 나누려면 Create부터 시작한다.
	f, err := os.Create("/tmp/dat2")
	check(err)
	
	defer f.Close()
	
	// byte slice를 Write 할수도 있음
	d2 := []byte{115, 111, 109, 101, 10}
	
	// Write()함수는 쓴 글자의 byte수, err을 리턴
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	
	// Sync() 함수를 이용, 안정된 보관소에 쓴 걸 flush한다.
	f.Sync()
	
	// bufio는 buffered reader에 더해 buffered writer를 제공
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	
	// Flush로 모든 buffered operation이 적용되는 것을 보증한다.
	w.Flush()
}