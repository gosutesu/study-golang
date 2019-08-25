package main

import "fmt"

// 構造体
type Vertex struct {
	X int
	Y int
}

/**
 * 構造体のポインタ
 */
func main() {
	v := Vertex{1, 2}
	p := &v

	p.X = 1e9

	fmt.Println(v)
}
