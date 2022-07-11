package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

func iAbs(a int) int {
	return int(math.Abs(float64(a)))
}

func l1Distance(c1, c2 Coordinate) int {
	return iAbs(c1.x-c2.x) + iAbs(c1.y-c2.y)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func readLine(sc *bufio.Scanner) (int, Coordinate) {
	t := nextInt(sc)
	x := nextInt(sc)
	y := nextInt(sc)
	return t, Coordinate{x, y}
}

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	t := 0
	c := Coordinate{0, 0}
	tPrev := 0
	cPrev := Coordinate{0, 0}
	for i := 0; i < N; i++ {
		t, c = readLine(sc)
		tNow := t - tPrev
		distance := l1Distance(c, cPrev)

		if tNow < distance {
			// 座標に届かない
			fmt.Println("No")
			return
		}

		if tNow%2 != distance%2 {
			// 座標に届くが、タイミングよく止まれない
			fmt.Println("No")
			return
		}

		cPrev = c
		tPrev = t
	}

	fmt.Println("Yes")
}
