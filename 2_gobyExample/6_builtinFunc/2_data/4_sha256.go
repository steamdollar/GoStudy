package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	
	// start new hash
	h := sha256.New()
	
	// Write()는 bytes를 받는다. string을 hash화 하고 싶다면 byte로 바꿔줘야함
	h.Write([]byte(s))
	
	// get finalizes hash result ad byte slice
	// Sum()의 param으론 보통 nil이 들어감
	bs := h.Sum()
	
	fmt.Printf("%x\n", bs)
}