package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	pattern := 0
	i := a
	for i >= 0 {
		if i*500 > x {
			i--
			continue
		}
		j := b
		for j >= 0 {
			if i*500+j*100 > x {
				j--
				continue
			}
			k := c
			for k >= 0 {
				// fmt.Println(i, j, k, pattern)
				if i*500+j*100+k*50 > x {
					k--
					continue
				}
				if i*500+j*100+k*50 == x {
					pattern++
					break
				}
				k--
			}
			j--
		}
		i--
	}

	fmt.Println(pattern)
}
