package main

// usr/lib/go/src

import "fmt"

import "calculation"

func main() {
	x,y := 15, 10

	sum := calculation.Do_add(x,y)
	fmt.Println(sum)

	sub := calculation.Do_sub(x,y)
	fmt.Println(sub)
}