package main

import "fmt"

// sliceのゼロ値はnilという。
// 0の長さと容量を持っており、何の元となる配列も持っていません。
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
