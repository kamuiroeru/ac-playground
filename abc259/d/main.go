package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

type Coordinate struct {
	x, y float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%.6f, %.6f)", c.x, c.y)
}

type Circle struct {
	xy Coordinate
	r  float64
}

func (c Circle) String() string {
	return fmt.Sprintf("%v, %.6f", c.xy, c.r)
}

func (c Circle) Across(coordinate Coordinate) bool {
	return math.Pow(coordinate.x-c.xy.x, 2)+math.Pow(coordinate.y-c.xy.y, 2) == math.Pow(c.r, 2)
}

func readLine(sc *bufio.Scanner) Circle {
	sc.Scan()
	line := sc.Text()
	splited := strings.Split(line, " ")
	coordinate := Coordinate{
		float64(atoi(splited[0])),
		float64(atoi(splited[1])),
	}
	return Circle{
		coordinate,
		float64(atoi(splited[2])),
	}
}

func distance(c1, c2 Coordinate) float64 {
	return math.Sqrt(math.Pow(c1.x-c2.x, 2) + math.Pow(c1.y-c2.y, 2))
}

func intercept(c1, c2 Circle) bool {
	var cLarge, cSmall Circle
	if c1.r > c2.r {
		cLarge = c1
		cSmall = c2
	} else {
		cLarge = c2
		cSmall = c1
	}
	if math.Pow((cSmall.xy.x-cLarge.xy.x), 2)+math.Pow((cSmall.xy.y-cLarge.xy.y), 2) < math.Pow(cLarge.r, 2) {
		// 同心円状になっている時
		if cLarge.r-cSmall.r > distance(cLarge.xy, cSmall.xy) {
			return false
		} else {
			return true
		}
	} else {
		// ベン図の時
		if distance(cLarge.xy, cSmall.xy) > cLarge.r+cSmall.r {
			return false
		} else {
			return true
		}
	}
}

func combination2(circles []Circle) [][]Circle {
	outer := make([][]Circle, 0, len(circles))
	for i := 0; i < len(circles)-1; i++ {
		for j := i + 1; j < len(circles); j++ {
			outer = append(outer, []Circle{circles[i], circles[j]})
		}
	}
	return outer
}

func combination(n int) [][]int {
	baseRange := make([]int, n)
	for i := range baseRange {
		baseRange[i] = i
	}
	outer := make([][]int, 0, n*(n-1)/2)
	for i := 0; i < len(baseRange)-1; i++ {
		for j := i + 1; j < len(baseRange); j++ {
			outer = append(outer, []int{baseRange[i], baseRange[j]})
		}
	}
	return outer
}

func aIndexList(a int, list []int) int {
	for i, elem := range list {
		if elem == a {
			return i
		}
	}
	return -1
}

func aInList(a int, list []int) bool {
	if aIndexList(a, list) >= 0 {
		return true
	} else {
		return false
	}
}

func aRemoveFromList(a int, list *[]int) {
	index := aIndexList(a, *list)
	if index >= 0 {
		*list = append((*list)[:index], (*list)[index+1:]...)
	}
}

func main() {
	var N, sx, sy, tx, ty int
	fmt.Scan(&N)
	fmt.Scanf("%d %d %d %d\n", &sx, &sy, &tx, &ty)
	s := Coordinate{float64(sx), float64(sy)}
	t := Coordinate{float64(tx), float64(ty)}

	sc := bufio.NewScanner(os.Stdin)
	circles := make([]Circle, 0, N)
	for i := 0; i < N; i++ {
		circles = append(circles, readLine(sc))
	}

	// fmt.Println(N)
	// fmt.Println(s)
	// fmt.Println(t)
	// fmt.Println(circles)

	// s と t がどの円周上にあるか見つける
	acrossIndexS := -1
	acrossIndexT := -1
	for i := 0; acrossIndexS < 0 || acrossIndexT < 0; i++ {
		c := circles[i]
		if c.Across(s) {
			acrossIndexS = i
		}
		if c.Across(t) {
			acrossIndexT = i
		}
	}

	if acrossIndexS == acrossIndexT {
		fmt.Println("Yes")
		return
	}

	// fmt.Println(acrossIndexS)
	// fmt.Println(acrossIndexT)

	canMove := make(map[int][]int, len(circles))
	for _, ij := range combination(len(circles)) {
		i := ij[0]
		j := ij[1]
		if intercept(circles[i], circles[j]) {
			canMove[i] = append(canMove[i], j)
			canMove[j] = append(canMove[j], i)
		}
	}

	// fmt.Println(canMove)

	footprint := make([]int, 0)
	nextMove := make([]int, 0)
	footprint = append(footprint, acrossIndexS)
	nextMove = append(nextMove, canMove[acrossIndexS]...)
	for len(nextMove) > 0 {
		for _, i := range nextMove {
			// fmt.Println(lc, "fp", footprint)
			// fmt.Println(lc, "nm", nextMove)

			if !aInList(i, footprint) {
				footprint = append(footprint, i)
			}

			if aInList(acrossIndexT, footprint) {
				fmt.Println("Yes")
				return
			}
			for _, j := range canMove[i] {
				if aInList(j, footprint) {
					// fmt.Println("removeFromList, j:", j)
					aRemoveFromList(j, &nextMove)
				} else {
					footprint = append(footprint, j)
					nextMove = append(nextMove, j)
				}
			}
		}
	}

	fmt.Println("No")
}
