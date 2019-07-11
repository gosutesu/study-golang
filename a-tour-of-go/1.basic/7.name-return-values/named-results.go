package main

import "fmt"

// 戻り値になる変数に名前をつける
// 戻り値に名前をつけると関数内で定義した変数名と扱われる。
// 戻り値に名前をつけることで、returnステートメントに何も書かずに戻ることができる。(naked return)
// しかし、可読性への影響があるため短い関数のみで使用すべき。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
