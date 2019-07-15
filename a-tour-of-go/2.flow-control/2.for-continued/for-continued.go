package main

import "fmt"

// for文について
// 初期化ステートメントと後処理ステートメントの省略が可能
// 他の言語でのwhile文はgoではfor文を使用する。
func main() {
	sum := 1
	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)
}
