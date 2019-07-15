package main

import "fmt"

// for文について
// for文は()を書く必要ない
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
