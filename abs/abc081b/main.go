package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Board []int

// 1行読み込んで Board にする
func readLineToIntArr() Board {
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

// Board の数がすべて 偶数 かどうか返す
func (b Board) isAllEven() bool {
	for _, i := range b {
		if i%2 == 1 {
			return false
		}
	}
	return true
}

// Board の中身をすべて 2 で割る
func (b *Board) divide2() {
	for i := 0; i < len(*b); i++ {
		(*b)[i] /= 2
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	board := readLineToIntArr()
	// fmt.Println(n)
	// fmt.Println(board)

	count := 0
	for board.isAllEven() {
		board.divide2()
		count++
	}

	fmt.Println(count)
}
