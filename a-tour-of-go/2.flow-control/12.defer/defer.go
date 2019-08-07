package main

import "fmt"

// defer句について
// 呼び出し元の関数が終了するまで処理を遅延させる処理
// 評価自体はすぐに実行されるが、returnがされない
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
