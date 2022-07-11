package main

import (
	"fmt"
)

type Store []string

func (s Store) includeEmpty() bool {
	for _, str := range s {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	keywords := []string{
		"dream",
		"dreamer",
		"erase",
		"eraser",
	}
	remains := make(Store, 0, 32)
	remains = append(remains, s)

	for {
		// pop
		remain := remains[len(remains)-1]
		remains = remains[:len(remains)-1]

		for _, kw := range keywords {
			kwLen := len(kw)
			if len(remain) < kwLen { // keyword よりも残りが短い時はスキップ
				continue
			}
			prefix := remain[:kwLen]
			suffix := remain[kwLen:]
			if prefix == kw {
				remains = append(remains, suffix)
			}
		}

		// prefix がどの keyword とも一致しなかった
		if len(remains) == 0 {
			fmt.Println("NO")
			break
		}

		// prefix がいずれかの keyword と一致し、suffix が無い（ピッタリ）
		if remains.includeEmpty() {
			fmt.Println("YES")
			break
		}
	}

}
