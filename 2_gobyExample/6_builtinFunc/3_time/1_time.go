package main

import (
	"fmt"
	"time"
)

var f = fmt.Println

/*
 	go에서 제공하는 다양한 시간, 기간 관련 기능들
*/

func main() {
	now := time.Now()
	f(now)
	
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		// y, m, d, h, m, s, ns
		
	f(then)
	// f(then.Year())
    // f(then.Month())
    // f(then.Day())
    // f(then.Hour())
    // f(then.Minute())
    // f(then.Second())
    // f(then.Nanosecond())
    // f(then.Location())
	
	// 요일
	f(then.Weekday())

	
	// then과 now를 비교 - return t/f
	//f(then.Before(now))
    //f(then.After(now))
    //f(then.Equal(now))
	
	// now에서 then까지의 시간차
	diff := now.Sub(then)
	f(diff)

	// 단위 convert
	f(diff.Hours())
	f(diff.Minutes())
	f(diff.Seconds())
	f(diff.Nanoseconds())
	
	// 시간의 +,-
	f(then.Add(diff))
	f(then.Add(-diff))
	
	
	//
	// unix epoch로 변환
	f(now.Unix())
	f(now.UnixMilli())
    f(now.UnixNano())
	
	// 반대로 unix epoch를 시간 형식으로 바꾸는 것도 가능
	f(time.Unix(now.Unix(), 0))
	f(time.Unix(0, now.UnixNano()))
}