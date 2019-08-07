package main

import (
	"fmt"
	"time"
)

// switchぶん
// 条件を指定しなくても書けるみたい。用途が思い浮かばず、、、無念
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good mornig!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Goot evening")
	}
}
