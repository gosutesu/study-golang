package main

import "fmt"

// 変数宣言の省略について
// 関数の中ならば、varを省略して変数宣言を行うことができる。
// varの代わりに := を使用して代入を行う。
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
