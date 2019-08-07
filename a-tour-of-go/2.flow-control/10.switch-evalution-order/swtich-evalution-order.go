package main

import (
	"fmt"
	"time"
)

// switch文について
// 書き方は他の言語と対して変わらないが、breakがいらないのが新鮮
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Int two days")
	default:
		fmt.Println("Too far away.")
	}
}
