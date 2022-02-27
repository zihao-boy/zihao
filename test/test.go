package main

import (
	"fmt"
	"math"
)

func main() {
	var count int64 = 10000

	a := float64(count)/1000
	fmt.Println(a)
	fmt.Println(math.Ceil(a))
}