package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*
		선행 명령어
		
		echo "hello" > /tmp/dat
		echo "go" >> /tmp/dat
		
		이 파일을 만들어야 읽던 말던 가능하다.
	*/
	
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	
	// ReadFile()은 param을 []byte로 읽어 옴
	// fmt.Println(dat)
	fmt.Println(string(dat))
	
	// file open > 어떤 파일을 열지, 파일의 어떤 부분을 조작할지를 정하는 첫 단계
	f, err := os.Open("/tmp/dat")
	check(err)
	
	defer f.Close()
	
	// 파일의 시작부터 5바이트를 읽겠다.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
	// Seek > 파일의 특정 위치를 지정해 거기서부터 읽을 수 있음
	o2, err := f.Seek(6,0)
	check(err)
	// fmt.Println(o2) // 6
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Println()
	fmt.Printf("%v\n", string(b2[:n2]))
	
	// io package는 읽기에 유용한 몇 가지 함수들을 제공한다.
	o3, err := f.Seek(6,0)
	check(err)
	
	b3 := make([]byte, 2)
	// e.g. ReadAtLeast 함수로 읽는 최소 글자를 지정 할 수도 있음
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	
	_, err = f.Seek(0,0)
	check(err)
	
	// bufio도 읽기 기능을 지원함
	// 여러 개 적은 read를 수행할 때 효율이 좋다.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	
}