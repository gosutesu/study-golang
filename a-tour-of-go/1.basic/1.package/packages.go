// Goのプログラムはパッケージで構成されている。
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favarite number is", rand.Intn(10))
}
