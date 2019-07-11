package main

import "fmt"

// functions について
// 同じ型の引数が２つ以上の場合、型を省略して書くことが可能です。
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(43, 12))
}
