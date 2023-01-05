package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	디렉토리 관련 함수들도 많다.
*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 현재 위치에서 새로운 폴더 생성
	err := os.Mkdir("subdir", 0755)
	check(err)
	
	// RemoveAll 은 rm -rf랑 비슷한 기능
	defer os.RemoveAll("subdir")
	
	// 새 빈 파일 생성
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	
	createEmptyFile("subdir/file1")
	
	// MkdirAll로 여러 depth의 경로, 파일을 생성할 수 있음
	// mkdir -p 와 동일
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	
	// ReadDir은 해당디렉토리 컨텐츠를 리스팅
	c, err := os.ReadDir("subdir/parent")
	check(err)
	
	fmt.Println("listing subdir/parent")
	// 리스팅한 (ls) 값을 출력
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	
	// 현재 디렉토리를 바꿀 수 있음 (cd)
	err = os.Chdir("subdir/parent/child")
	check(err)
	
	fmt.Println("listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	
	// 원래 경로로 돌아옴
	err = os.Chdir("../../..")
	check(err)
	
	// 디렉토리를 모든 서브 디렉토리에 대해 재귀적으로 들어갈 수도 있다.
	// Walk() 함수를 이용해 콜백 함수로 각각에 대해 실행하고 싶은 함수 호출 가능
	fmt.Println("visiting subdir")
	// 현 경로의 모든 서브경로에 대해 visit 함수 실행
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}