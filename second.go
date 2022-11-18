/*

package 

코드를 조직화할 때 사용함. 

import package_name

// go를 설치한 폴더에 이 파일을 넣어야 가져다 쓸 수 있게 된다.

echo $PATH < 여기 없음 시발

/usr/lib/go/src

GO lang을 설치한 소스 폴더 경로가 필요하다.

거기 들어가서 Go 파일을 작성하고,

그걸 install 해야 다른 곳에서 불러와 쓸 수 있음

sudo go install

*/

package calculation

func Do_add(num1 int, num2 int) (int) {
	sum := num1 + num2
	return sum
}