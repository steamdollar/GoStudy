package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
	filepath 패키지가 file path를 만들고, parse하는걸 담당함
*/

var fl = fmt.Println

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fl("p:", p)
	
	// Join method를 사용해 path를 만드는게 좋다.
	// join은 표준화 기능까지 가지고 있어서 알아서 오타 등을 보정해줌
	fl(filepath.Join("dir1//", "filename"))
	fl(filepath.Join("dir1/../dir1", "filename"))
	
	// 이렇게 만든 path를 Dir(), Base()로 parse해 디렉토리, 이름을 따로 읽을 수도 있음
	fl("Dir(p) : ", filepath.Dir(p))
	fl("Base(p) : ", filepath.Base(p))
	
	// 절대 경로 여부를 판별도 해줌
	fl(filepath.IsAbs("dir/file")) // false
	fl(filepath.IsAbs("/dir/file")) // true
	
	//
	
	filename := "config.json"
	
	// Ext 함수를 이용해 확장자만 읽을 수 있음
	ext := filepath.Ext(filename)
	fl(ext) //.json
	
	// strings.TrimSuffix를 이용해 확장자 없이 파일 이름만 읽을 수 있음
	fl(strings.TrimSuffix(filename, ext)) // config
	
	// Rel()은 base와 target 사이의 상대 경로를 알려줌
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	// 상대 경로를 만들 수없다면 에러
	if err != nil {
		panic(err)
	}
	fl(rel)
	
	rel2, err := filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fl(rel2)
}