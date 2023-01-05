package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	잠깐 쓰고 버릴 tmp 파일들 다루기
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// temp file을 생성하는 가장 쉬운 방법은 CreateTemp를 사용하는 것
	// create, open까지 해줌 (for reading, writing)
	
	// 1st param은 경로를 의미, 공란으로 주면 OS의 디폴트 디렉토리에 생성
	// 2nd param은 temp file 이름의 prefix
	f, err := os.CreateTemp("", "sample")
	check(err)
	
	// 파일 이름 불러오기 (prefix + random Num)
	fmt.Println("temp file name:", f.Name())
	
	defer os.Remove(f.Name())
	
	// write sth to file
	_, err = f.Write([]byte{1,2,3,4})
	check(err)
	
	// 혹시 여러 임시 파일을 생성하고 싶으면 임시 경로를 생성하는게 좋다.
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("temp dir name", dname)
	
	defer os.RemoveAll(dname)
	
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1,2}, 0666)
	check(err)
}