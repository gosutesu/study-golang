package main

import "fmt"

type Vertrix struct {
	X, Y int
}

// 構造体の宣言の仕方によって初期値が違う
var (
	v1 = Vertrix{1, 2}
	v2 = Vertrix{X: 1}
	v3 = Vertrix{}
	p  = &Vertrix{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
