package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 構造体のメンバへのアクセスの仕方
func main() {
	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)
}
