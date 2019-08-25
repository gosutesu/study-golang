package main

import "fmt"

// Sliceには長さと容量がある。
// 長さは要素の数である。
// 容量はスライスの最初の要素から考えて元となる配列の要素数である。
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	prinntSlice(s)

	s = s[:0]
	prinntSlice(s)

	s = s[:4]
	prinntSlice(s)

	s = s[2:]
	prinntSlice(s)
}

func prinntSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
