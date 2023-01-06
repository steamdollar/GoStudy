package main

import (
	"fmt"
	"flag"
	"os"
)

/*
	자체 flag set을 가진 subcommand를 가지는 cli tools이 있다. (git, go tool 등..)
	e.g. go tool의 subcommand는 go build go get이 있음
	
	이런 subcommand의 자체 flag도 마찬가지로 flag package를 이용해 설정할 수 있음
*/

var fl = fmt.Println

func main() {
	// NewFlagSet을 이용해 subcommand를 선언
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 1st pa : subcommand name, 2nd
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")
	
	// subcommand마다 다른 flag를 설정할 수 있음
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	
	// subcommnad가 cli의 1st arg로 와야함
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	
	switch os.Args[1] {
		// 각각의 subcommand에 대한 case문 
	case "foo" :
		fooCmd.Parse(os.Args[2:])
		fl("subcommand 'foo'")
		fl(" enable :", *fooEnable)
		fl(" name :", *fooName)
		fl(" tail :", fooCmd.Args())
		
	case "bar" :
		barCmd.Parse(os.Args[2:])
		fl("subcommand 'bar'")
		fl(" level : ", *barLevel)
		fl(" tail : ", barCmd.Args())
	
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
