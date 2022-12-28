package main
import (
	"fmt"
	"sort"
)

/*
	데이터 정렬 기능을 sort 패키지를 이용해 할 수 있다.
*/

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	/*
		sort하면 새로운 slice를 리턴하는게 아니라
		원래 slice를 변경해 돌려준다.
	*/
	
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	
	s := sort.IntsAreSorted(ints)
	// 이미 sort된 상태라면 true, vice versa
	fmt.Println("Sorted: ", s) 
}