package main
import (
	"os"
	"text/template"
	// "fmt"
)

/*
	text/template 패키지 : support for creating dynamic content 
	or showing customized output to user
	
	비슷한 html/template도 동일 API를 제공하지만 추가로 보안 측면에서 몇 가지 기능이
	더 있고, html 생성시 활용됨
*/

func main () {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	/* 
		새로운 템플릿을 만들고 그 바디를 string으로부터 parse할 수 있다.
		템플릿은 static text와 action의 혼합체가 {{...}} 안에 포함된 형태
	 	이 템플릿은 다양한 컨텐츠에 삽입되어 들어갈 수 있다.
	*/
	
	t1 = template.Must(t1.Parse("Value : {{.}}\n"))
	/* 
		다른 방법으로 template.Must 함수를 사용해 parse 결과가 error를 리턴할 경우
		panic을 일으키는 방법도 있다.
		global scope로 초기화되는 template에 대해 유용하다.
		fmt.Println(t1)
		&{t1 0xc00010c120 0xc000072050  }
	*/
	
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
	/* 
		template을 execute 함으로써 특정 값을 가진 text를 생성한다.
	 	{{.}} action이 execute의 파라미터로 전달되는 값으로 대체된다.
	*/
	
	Create := func(name, t string) *template.Template{
		return template.Must(template.New(name).Parse(t))
	}
	
	t2 := Create("t2", "Name : {{.Name}}\n")
	/* 
		data가 struct라면 {{.fieldName}} action을 사용해 필드에 접근할 수 있다.
		template가 execute 되는 중엔 필드가 export되어 접근 가능해진다.
	*/
	
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Alpha Beta"})
	
	t2.Execute(os.Stdout, map[string]string{
		"Name" : "Elicia",
	})
	// map도 동일하게 적용
	
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	/*
		if/else도 template에 대해 사용 가능
		0, nil 같은건 false로 간주, "-"는 공백 제거를 위해 사용
	*/
	
	t4 := Create("t4",
		"Range : {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"JAVA",	
		})
}