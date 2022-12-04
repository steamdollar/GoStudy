package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./txt/file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f)
	

	l1, err := f.WriteString("first line of the file \n")

	l2, err := f.WriteString("한글도 되나2?")
	// l은 length를 의미

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l1, "bytes written")
	fmt.Println(l2, "bytes written")
	// err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}