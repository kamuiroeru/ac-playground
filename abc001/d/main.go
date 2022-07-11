package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type RainRange struct {
	s int
	e int
}

func (rr RainRange) String() string {
	return fmt.Sprintf("%04d-%04d", rr.s, rr.e)
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func readLine(sc *bufio.Scanner) RainRange {
	sc.Scan()
	line := sc.Text()
	splited := strings.Split(line, "-")
	return RainRange{
		atoi(splited[0]),
		atoi(splited[1]),
	}
}

// 5分に丸める
func (rr *RainRange) fit() {
	rr.s = rr.s - (rr.s % 5)
	if rr.e%5 != 0 {
		rr.e = rr.e + 5 - (rr.e % 5)
	}

	// :60 の場合は繰り上げる。 Ex. 13:60 は 14:00 にしないとだめ
	rreStr := strings.Split(rr.String(), "-")[1]
	if rreStr[2:4] == "60" {
		rr.e += 40 // :60 の場合は繰り上げる 13:60 は 14:00 にしないとだめ
	}
}

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	rrs := make([]RainRange, 0, 32768)
	for i := 0; i < N; i++ {
		rr := readLine(sc)
		rr.fit()
		rrs = append(rrs, rr)
	}

	sort.Slice(rrs, func(i, j int) bool {
		return rrs[i].s < rrs[j].s
	})

	mergedRrs := make([]RainRange, 0, 32768)
	mergedRrs = append(mergedRrs, rrs[0])
	for i := 1; i < len(rrs); i++ {
		rr := rrs[i]
		tail := &mergedRrs[len(mergedRrs)-1]
		if rr.s <= tail.e {
			if tail.e <= rr.e {
				tail.e = rr.e
			}
		} else {
			mergedRrs = append(mergedRrs, rr)
		}
	}

	for _, rr := range mergedRrs {
		fmt.Println(rr)
	}
}
