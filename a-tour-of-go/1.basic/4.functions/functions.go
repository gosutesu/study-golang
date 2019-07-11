package main

import "fmt"

// functions について
// 型は変数の後ろにつける。
// なぜかは「https://blog.golang.org/gos-declaration-syntax」を参照する
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
