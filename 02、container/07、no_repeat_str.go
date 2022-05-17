package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
func main() {
	// 寻找最长不含有重复字符的字串
	// abcadbcbb -> abcd
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcdbb"))
	fmt.Println(lengthOfNonRepeatingSubstr("一二三二一"))
}
