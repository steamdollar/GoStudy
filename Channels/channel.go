/*

channel은 함수 간에 상호작용을 구현하는 방법이다. 

한 routine이 데이터를 올리면, 다른 routine이 이에 접근할수 있도록 

그 매개체가 된다고 생각하면 된다.

채널 선언 syntax는 다음과 같다.


chan_name := make(chan datatype)

get_info := make(chan string)

ch := make(chan int)


채널로 데이터를 전송하는 코드는 다음 syntax를 따른다.

chan_name <- var_name

chan_name에 var_name을 넣는다.

ch <- x

x를 채널로 전송

채널로부터 데이터를 받는 것은 다음 syntax를 따른다.

var_name := <- chan_name

y := <- ch

ch에서 받은 데이터를 y에 넣는다. 그래서 변수 선언 기호가 있는 듯

//

앞서 goroutine을 공부할 때, main program이 goroutine을 기다리지 않는다고 배웠는데,

channel이 연관되어 있다면 달라진다. 

goroutine이 channel에 데이터를 넣으면 main() 함수는 channel이 데이터를 받을 때 까지 기다린다.

*/

package main
import "fmt"
// import "time"

func display(ch chan int) {
	// time.Sleep ( 5 * time.Second )
	fmt.Println("Inside display()")
	ch <- 1234
}

func main() {
	cha := make(chan int)
	go display(cha)
	x := <- cha
	fmt.Println("Inside main()")
	fmt.Println("Printing x main() after taking from channel: ", x)
}

/*

x := <- ch 가 실행되려면 channel의 data를 먼저 수신해야한다.

즉, ch 를 받을때까지 기다려야 한다.

display() 함수는 5초를 기다렸다가 실행되고 데이터를 channel ch로 push한다.

main()이 데이터를 channel로부터 데이터를 수신하면 unblock되어 계속 실행된다.

Inside display()
Inside main()
Printing x main() after taking from channel:  1234

*/