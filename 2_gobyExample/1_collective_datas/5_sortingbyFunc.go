package main
import (
	"fmt"
	"sort"
)

/*
	디폴트가 아닌 다른 규칙에 따라 정렬하고 싶다면..
*/

type byLength []string
/*
	Go의 커스텀 함수를 이용해 sort하려면 대응하는 type이 필요하다.
	여기서는 []string type에 대응ㅎ라는 byLength type을 만듦
*/

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}
/*
	sort.interface - Len, less, swap을 implement
	> sort 패키지의 generic Sort 함수를 사용할 수 있다.
	여기선 less만이 특정 로직에 따른 sort (나머지 둘은 타입 무관 보편적임)
	(글자 길이에 따른 정렬)
*/

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
// 이게 furits slice를 byLength로 convert해 custom sort method를 implement 하고
// sort.Sort 를 typed slice에 사용할 수 있다. 