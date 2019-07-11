package main

import (
	"fmt"
	"math"
)

// export について
// Goでは最初の文字が大文字で始まる名前が外部のパッケージに公開されている名前です。
// 小文字で公開されているものは公開されていない名前です。
func main() {
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
