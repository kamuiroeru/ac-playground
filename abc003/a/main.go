package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	sum := 0
	for i := 1; i <= N; i++ {
		sum += i * 10000
	}

	fmt.Println(sum / N)
}
