package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) != 0 {
			if strings.Index(strs[i], prefix) == 0 {
				break
			} else {
				prefix = prefix[:len(prefix)-1]
			}
		}
	}
	return string(prefix)
}
func main() {
	strs := []string{"kaka", "kaka1", "kaka2"}
	fmt.Println(strs[1])
	prefix := longestCommonPrefix(strs)
	fmt.Println("prefix = ", prefix)
}
