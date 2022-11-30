package main

import (
	"fmt"
	"os"
	"errors"
)

/*
	직접 에러 메시지를 만들고 싶다면 error 패키지의 New() 함수를 사용하면 된다.
	 
	

*/

func fileopen(name string) (string, error) {
	f, er := os.Open(name)

	if er != nil {
		return "", errors.New("Custom error Message: File name is wrong")
	} else {
		return f.Name(), nil
	}
}

func main() {
	filename, error := fileopen("invalid.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("file opened", filename)
	}
}