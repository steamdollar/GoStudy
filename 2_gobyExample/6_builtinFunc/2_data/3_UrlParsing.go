package main

import (
	"fmt"
	"net"
	"net/url"
)

var fl = fmt.Println

func echoline() {
	fl("=============")
}


func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v&k2=v2#f"
	
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	
	fl(u.Scheme)
	
	echoline()
	
	fl(u.User)
	fl(u.User.Username())
	p, _ := u.User.Password()
	fl(p)
	
	echoline()
	
	fl(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fl(host)
	fl(port)
	
	fl(u.Path)
	fl(u.Fragment)
	// url의 #뒤에 있는게 fragment
	
	fl(u.RawQuery)
	// ? 뒤에 있는거
	
	m, _ := url.ParseQuery(u.RawQuery)
	// query param들을 k-v 형태로 묶고 싶다면 이렇게
	
	fl(m)
	fl(m["k2"][0])
}