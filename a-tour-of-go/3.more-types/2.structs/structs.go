package main

import "fmt"

// 構造体自体
type Vertrix struct {
	X int
	Y int
}

// 構造体の宣言の仕方を
func main() {
	fmt.Println(Vertrix{1, 2})
}
