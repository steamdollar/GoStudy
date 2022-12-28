package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id, attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

var f = fmt.Println

/*
	plant struct가 xml에 매핑된다.
	json 예시와 비슷하게 인코더, 디코더를 위한 디렉티브가 field tag에 포함된다.
	
	여기서는 XML package의 특성을 활용하는데, 
	XMLName field가 이 struct를 대표하는 xml element의 이름을 알려줌
	id, attr은 id field가 nested element가 아닌 XML attribute임을 지정
*/

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", 
	p.Id, p.Name, p.Origin)
}

func main () {
	coffee := &Plant{Id : 27, Name : "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	// f(string(out))
	// plant를 represent하는 XML을 emit
	// MarshalIndent를 이용해 사람이 읽을 수 있는 output 생성
	f(xml.Header + string(out))
	// header를 추가하려면 명시적으로 append 해준다.
	
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	f(p)
}