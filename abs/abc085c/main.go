package main

import (
	"fmt"
)

func main() {
	var N, Y int
	fmt.Scan(&N)
	fmt.Scan(&Y)

	// 3元2連立方程式なので、 x を定数だと考えて(2001回)、2元2連立方程式を解く
	for x := 0; x < 2000+1; x++ {
		fourY := Y/1000 - N - 9*x
		fourZ := -Y/1000 + 5*N + 5*x
		// fmt.Println(fourY, fourZ)
		if fourY == 0 {
			fmt.Println(x, 0, N-x)
			return
		}
		if fourZ == 0 {
			fmt.Println(x, N-x, 0)
			return
		}
		if fourY > 0 && fourZ > 0 {
			if fourY%4 == 0 && fourZ%4 == 0 {
				fmt.Println(x, fourY/4, fourZ/4)
				return
			}
		}
	}

	fmt.Println(-1, -1, -1)
}
