package main

import (
	"encoding/xml"
	// 이 패키지를 이용해 xml 계열 포맷을 지원한다.
	"fmt"
)

var f = fmt.Println

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id, attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

/*
	plant struct가 xml에 매핑된다.
	json 예시와 비슷하게 field tag는 인코더, 디코더를 위한 디렉티브를 포함한다.
	
	여기서는 XML package의 특성을 활용하는데, 
	XMLName field가 이 struct를 대표하는 xml element의 이름을 알려줌
	id, attr은 id field가 nested element가 아닌 XML attribute임을 지정
*/

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", 
		p.Id, p.Name, p.Origin)
}

func main () {
	coffee := &Plant{Id : 27, Name : "coffee", Origin : []string{"Ethiopia", "Brazil", "Colombia"} }
	
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	// f(string(out))
	// plant를 represent하는 XML을 emit
	// MarshalIndent를 이용해 사람이 읽을 수 있는 output 생성
	// f(xml.Header + string(out))
	// header를 output에 추가하려면 명시적으로 append 해준다.
	// 메타데이터 담는 그 header 맞음
	
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	// f(p)
	// Unmarshal method을 이용해 XML 데이터를 parse한다.
	
	//
	tomato := &Plant{Id : 81, Name : "Tomato"}
	tomato.Origin = []string{"Mexico", "Califonia"}
	
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
		// 모든 plant를 <parent><child> 하위에 두겠다는 의미 (띄어쓰기 없음)
	}
	
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	f(string(out))
}