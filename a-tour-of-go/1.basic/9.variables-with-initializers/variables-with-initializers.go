package main

import "fmt"

var i, j int = 1, 2

// 変数宣言の初期化について
// 初期化子が与えられている場合、型を省略可能
func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
