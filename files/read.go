package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")
	
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}

	fmt.Println("Contents of file:\n", string(data))
}