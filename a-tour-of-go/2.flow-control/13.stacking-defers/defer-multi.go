package main

import "fmt"

// 複数のdeferについて
// deferはLIFOの順に実行される。
func main() {
	fmt.Println("couting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
