package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	aIsOdd := a%2 == 1
	bIsOdd := b%2 == 1
	result := "Even"
	if aIsOdd && bIsOdd {
		result = "Odd"
	}
	fmt.Println(result)
}
