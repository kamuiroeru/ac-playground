package main

import (
	"fmt"
	"strings"
)

// s からつながっている部分ごとに抜き出した string[] を返す
// SplitRepeat("abbaaccc") -> ["a", "bb", "aa", "ccc"]
func SplitRepeat(s string) []string {
	counter := make([]string, 0, 1024)
	temp := []string{s[0:1]}
	for i := 1; i < len(s); i++ {
		c := s[i : i+1]
		if c != temp[0] {
			counter = append(counter, strings.Join(temp, ""))
			temp = []string{c}
		} else {
			temp = append(temp, c)
		}
	}
	counter = append(counter, strings.Join(temp, ""))
	return counter
}

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)

	srS := SplitRepeat(S)
	srT := SplitRepeat(T)

	// 文字列の繋がりの個数が違う
	// Ex. S: aabbcc, T: aabb
	if len(srS) != len(srT) {
		fmt.Println("No")
		return
	}

	// 同じ文字の繰り返し文字列ごとに確認する
	for i := 0; i < len(srS); i++ {
		sS := srS[i]
		sT := srT[i]

		// 文字種類の並び方が違う
		// Ex. S: aaccbb, T: aabbcc
		if sS[0] != sT[0] {
			fmt.Println("No")
			return
		}

		// S 側の文字が 1文字 なのに、 Tは複数ある
		// Ex. S: abb, T: aaaabb
		if len(sS) == 1 && len(sT) > 1 {
			fmt.Println("No")
			return
		}

		// S 側の文字数 Tの文字数よりも多い
		// Ex. S: aaabb, T: aabb
		if len(sS) > len(sT) {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
