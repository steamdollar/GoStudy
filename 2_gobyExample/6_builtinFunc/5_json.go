package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/* json encoding, decoding in go */

var f = fmt.Println

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}
// export 된 field만이 JSON으로 변환, 역변환 가능하다.
// export 되려면 field 첫자를 대문자로 써야한다. 

func main() {
	bolB, _ := json.Marshal(true)
	f(string(bolB))
	
	intB, _ := json.Marshal(1)
	f(string(intB))
	
	fltB, _ := json.Marshal(2.34)
	f(string(fltB))
	
	strB, _ := json.Marshal("gopher")
	f(string(strB))
	// 기본 타입을 json string으로 인코딩
	
	slcD := []string{"apple", "lettuce" , "pear"}
	slcB, _ := json.Marshal(slcD)
	f(string(slcB))
	
	mapD := map[string]int{"apple" : 5, "lettuce" : 7}
	mapB, _ := json.Marshal(mapD)
	f(string(mapB))
	
	// map, slice를 json 객체로 인코딩
	
	res1D := &response1{
		Page : 1,
		Fruits : []string{"apple", "peach", "pear"}}
	
	res1B, _ := json.Marshal(res1D)
	f(string(res1B))
	// 커스텀 데이터 타입을 자동으로 인코딩해줄수 있음
	// 
	
	res2D := &response2 {
		Page : 1,
		Fruits : []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	f(string(res2B))
	// Json key 이름을 struct 선언시 지정해 줄 수 있다.
	
	// 이제 json을 go value로 디코딩을 해보자 \\
	
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)
	
	var dat map[string]interface{}
	// json package가 디코드된 데이터를 넣을 변수가 필요하다.
	
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	f(dat)
	// decode, check err
	
	num := dat["num"].(float64)
	f(num)
	// 디코딩 map의 value를 사용하기 위해 적절한 type으로 변환해주어야 한다.
	// 여기서 num의 value를 쓰려면 이걸 float64로 바꿔줘야 함
	
	strs := dat["strs"].([]interface{})
	// strs를 디코드
	str1 := strs[0].(string)
	// nested data는 여러 단계 변환을 거쳐야 함
	f(str1)
	
	// json을 커스텀 데이터 타입으로 디코드도 가능 \\
	
	str := `{"page":1, "fruits" : ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	// 이렇게 커스텀 타입으로 변환하면 타입 측면에서 안전성에 더해
	// 이 후 데이터 접근 시 타입 검증을 안해도 된다는 장점이 있음
	f(res)
	f(res.Fruits[0])
	
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple" : 5, "lettuce" : 7}
	enc.Encode(d)
	// 위의 예시들에서는 bytes와 string을 data와 json의 표준출력의 중간체로 사용했지만
	// json 인코딩 데이터를 바로 os.Writer나 os.Stdout 
	// 혹은 Http response body로도 직접 변환할 수 있다.
}