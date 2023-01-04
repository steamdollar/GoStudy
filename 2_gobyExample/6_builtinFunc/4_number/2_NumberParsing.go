package main

import (
	"fmt"
	"strconv"
	// strconv 패키지가 string > number parse를 담당
)

/*
	string > number parsinf 
*/

var fl = fmt.Println

func main() {
	
	// 소수의 convert, 2nd param은 몇 비트로 parse 할 것인지를 의미
	f, _ := strconv.ParseFloat("1.234", 64)
	fl(f)
		
	// 
	i, _ := strconv.ParseInt("123", 0, 64)
	fl(i)
	
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fl(d)
	
	u, _ := strconv.ParseUint("789", 0, 64)
	fl(u)
	
	k, _ := strconv.Atoi("135")
	fl(k)
	
	_, e := strconv.Atoi("wat")
	fl(e) 
}