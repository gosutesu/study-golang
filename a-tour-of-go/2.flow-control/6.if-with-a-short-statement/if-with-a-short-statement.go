package main

import (
	"fmt"
	"math"
)

// if文
// 評価式の前に簡単なステートメントを書くことが可能である。
// スコープはif文の中まで
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
