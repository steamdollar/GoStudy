package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	
	// b64 encoding, []byte 로 받아야함.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	
	// Decode는 에러가 날수도 있으므로 다음과 같이 확인할 수 있다.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	
	fmt.Println()
	
	// URlEncoding은 URL-compatible한 base64 format을 사용
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	// uEnc, sEnc가 약간 다르지만 디코드하면 원래 string을 잘 돌려줌
	
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}