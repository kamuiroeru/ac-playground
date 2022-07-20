package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	a := S[0:1]
	b := S[1:2]
	c := S[2:3]

	if a == b {
		if b == c {
			fmt.Println(-1)
		} else {
			fmt.Println(c)
		}
	} else {
		if a == c {
			fmt.Println(b)
		} else {
			fmt.Println(a)
		}
	}
}
