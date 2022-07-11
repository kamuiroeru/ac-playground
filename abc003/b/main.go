package main

import (
	"fmt"
	"strings"
)

const replaceable = "atcoder"

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)

	for i := 0; i < len(S); i++ {

		if S[i:i+1] == T[i:i+1] {
			continue
		}

		if S[i:i+1] == "@" {
			if !strings.Contains(replaceable, T[i:i+1]) {
				fmt.Println("You will lose")
				return
			}
			continue
		}

		if T[i:i+1] == "@" {
			if !strings.Contains(replaceable, S[i:i+1]) {
				fmt.Println("You will lose")
				return
			}
			continue
		}

		if S[i:i+1] != T[i:i+1] {
			fmt.Println("You will lose")
			return
		}
	}

	fmt.Println("You can win")
}
