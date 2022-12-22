package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}
// container struct는 base struct를 embed하고 있다.

type container2 struct {
	base
	dex int
}

func main() {
	co := container {
		base : base {
			num :1,
		},
		str : "some name",
	}
	// 리터럴로 struct를 생성할 때는 embed된 대상을 명시해주어야 한다.
	// 여기서 embedded type은 field name처럼 사용된다.

	fmt.Printf("co= {num : %v, str : %v}\n", co.num, co.str)
	// embedded 된 struct의 속성에도 바로 접근할 수 있다.

	fmt.Println("also num:", co.base.num)
	// 물론 full path를 적어도 된다.

	fmt.Println("describe:", co.describe())
	// base가 container에 embedded 되었다는 말은
	// base의 method가 곧 container의 method가 된다는 것을 의미한다.
	// 역시 full path를 쓰지 않고 container instance에서 바로 호출 가능

	type describer interface {
		describe() string
	}

	var d describer = co
	// 이게 왜 가능한가..
	// describer가 describe를 embed하고 있고,
	// 이 describe는 containter struct를 매개변수로 가지기 때문
	// 관계의 역전이라고 볼 수 있는건가..
	// describer가 embed한 describe는 매개변수를 container만 가지니까
	// 자동으로 go에선 describe method를 embed한 describer아 container type이라고 추리하는 것
	
	// 응 아냐 내일 다시 해봐..
	// base를 포함하는 type을 하나 더 만들어서 걔도 이런식으로 값 할당이 가능한지 확인해볼것
	// 가능하네..
	//

	co2 := container2 {
		base : base {
			num : 2,
		},
		dex : 3,
	}

	var d2 describer = co2


	fmt.Println("describer:", d.describe())
	// struct에 method를 embed 하면 다른 struct에 interface를 implement할 수 있다.
	// 예시에서 container가 describer interface를 implement 한다.
	// describer가 base를 embed하므로
	fmt.Println(d2.describe())

	// method를 embed한 struct가 있다면 이를 다른 struct에 implement할 수 있다.
	// describer는 base를 매개변수로 갖는 method describe를 embed 하는데,
	// 이러면 describer는 base를 embed한 다른 type struct의 form을 가질 수 있게 되는것
	// 이렇게 결론 내리는게 맞는 것 같다.
}