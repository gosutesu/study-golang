package main

import "fmt"

// 複数の戻り値
// GOは複数の戻り値を返却することが可能です。
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
