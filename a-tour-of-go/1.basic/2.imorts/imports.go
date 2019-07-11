package main

import (
	"fmt"
	"math"
)

// import に関しては下記のようにそれぞれ分けて書くことも可能。
// import "fmt"
// import "math"
// 上記のfactoredインポートの方がいいスタイルなので、上記の書き方に統一する。

func main() {
	fmt.Printf("Now you hava %g problems.", math.Sqrt(7))
}
