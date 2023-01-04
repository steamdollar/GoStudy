package main

import (
	"fmt"
	"time"
)

var p = fmt.Println
// := 선언은 이 영역에선 안되는듯

/*
	go는 패턴 기반 레이아웃을 통한 시간 formatting, parsing을 지원한다.
*/

func main() {
	
	// 시간을 RFC3339 format으로 변경
	t := time.Now()
	p(t.Format(time.RFC3339))
	
	// Format과 동일한 레이아웃을 사용한다.
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
    p(t1)
	
	p(t.Format("3:04PM"))
	p(t.Format("MON JAN _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)
	
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
	
	ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}