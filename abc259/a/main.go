package main

import (
	"fmt"
)

func main() {
	var N, M, X, T, D int
	fmt.Scanf("%d %d %d %d %d", &N, &M, &X, &T, &D)

	age0 := T - X*D

	if M < X {
		fmt.Println(D*M + age0)
		return
	}

	if X <= M && M <= N {
		fmt.Println(T)
		return
	}

}
