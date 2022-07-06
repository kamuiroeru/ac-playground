package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MochiSet map[int]struct{}

// 1行読み込んで MochiSet にする
func readLineToMochiSet(lineNumber int) MochiSet {
	ms := make(MochiSet, 100)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // 行区切りで読み込む

	for lc := 0; lc < lineNumber; lc++ {
		for scanner.Scan() {
			i, _ := strconv.Atoi(scanner.Text())
			ms[i] = struct{}{}
		}
	}
	return ms
}

func main() {
	var n int
	fmt.Scan(&n)
	ms := readLineToMochiSet(n)
	fmt.Println(len(ms))
}
