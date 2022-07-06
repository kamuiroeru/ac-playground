package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := 0
	for i := 1; i < n+1; i++ {
		digitSum := 0
		for _, digitStr := range strings.Split(strconv.Itoa(i), "") {
			digitInt, _ := strconv.Atoi(digitStr)
			digitSum += digitInt
		}
		if a <= digitSum && digitSum <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}
