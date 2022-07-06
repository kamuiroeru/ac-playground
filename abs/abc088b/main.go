package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Cards []int

// 1行読み込んで Cards にする
func readLineToIntArr() Cards {
	arr := make([]int, 0, int(math.Pow10(2)))
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // 行区切りで読み込む
	for scanner.Scan() {
		s = scanner.Text()
	}
	for _, iStr := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(iStr)
		arr = append(arr, i)
	}
	return arr
}

func (c *Cards) sortDesc() {
	sort.Slice(*c, func(i, j int) bool {
		return (*c)[i] > (*c)[j]
	})
}

func main() {
	var n int
	fmt.Scan(&n)
	cards := readLineToIntArr()
	cards.sortDesc()

	diff := 0
	for i := 0; i < len(cards); i++ {
		if i%2 == 0 {
			diff += cards[i]
		} else {
			diff -= cards[i]
		}
	}

	fmt.Println(diff)
}
