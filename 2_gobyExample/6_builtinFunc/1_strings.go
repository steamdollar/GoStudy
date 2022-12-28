package main

import (
	"fmt"
	s "strings"
	// strings 패키지를 s로 쓰겠다는 뜻
)

/*
	스탠다르 라이브러리 strings는
	문자열에 관한 함수들을 제공한다.
*/

var p = fmt.Println
// fmt.Println을 짧게 쓰기 위해 alias 한다.

func main () {
	p("Contains : ", s.Contains("test", "es"))
	// x가 y를 포함하는가? 
	
	p("Count : ", s.Count("test", "t"))
	// x에 y가 몇 개 포함되는지?
	
	p("HasPrefix : ", s.HasPrefix("test", "te"))
	// x가 y를 접두어로 받는가
	
	p("HasSuffix : ", s.HasSuffix("test", "est"))
	// x가 y를 접미어로 받는가
	
	p("Index : ", s.Index("test", "e"))
	// y가 x의 몇 번쨰 index인지 출력
	
	
	p("Join : ", s.Join([]string{"a", "b", "c"}, "-"))
	// string slice의 element들을 x를 사이에 두고 붙여줌
	
	p("Repeat : ", s.Repeat("ax", 5))
	// x를 y번 반복
	
	p("Replace : ", s.Replace("foo", "o", "q", 1))
	// x에서 y를 z로 앞에서부터 n개만 대체. n = -1이면 전부 대체
	
	p("Split : ", s.Split("a-b-c-d-e", "-") )
	// y를 기준으로 x를 split
	
	p("ToLower : ", s.ToLower("TEST"))
	// 문자열의 문자를 전부 소문자로
	
	p("ToUpper : ", s.ToUpper("test"))
	// 문자열의 문자를 전부 대문자로
}