package main

import (
	"embed"
)

/*
	go:embed는 프로그램이 임의의 파일, 폴더를 build시에 Go binary에 포함하도록 해주는 compiler directive이다.
	ref : https://pkg.go.dev/embed
*/

// embed directive는 Go source file을 포함한 디렉토리에 대한 상대 디렉토리를 받는다.
// 아래처럼 주석에 go:embed를 쓰고 가져올 파일의 상대경로를 입력해줘야함.

//go:embed folder/single_file.txt
var fileString string

//  혹은 file의 내용물을 byte로 받는다.
	
//go:embed folder/single_file.txt
var fileByte []byte

// 와일드카드를 이용해 여러 개 파일을 한번에 받거나, 폴더를 받는 것도 가능하다.
// 이때는 embed.FS type 변수를 사용한다.
// embed.FS는 simple virtual file system을 구현한다.

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main () {
	// single_file.txt의 내용물 출력
	print(fileString)
	print(string(fileByte))
	
	// embedded fioler에서 몇 개의 파일을 가져옴
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))
	
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}