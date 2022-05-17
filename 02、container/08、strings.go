package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "yes我爱慕课网a"
	fmt.Println(str)

	for _, ch := range []byte(str) { // UTF-8
		fmt.Printf("%X ", ch)
	}
	fmt.Println()

	for i, ch := range str { // ch is a rune 即 int32
		fmt.Printf("%d-%X ", i, ch) // UTF-8编码 转成了 Unicode编码
	}
	fmt.Println()

	bytes := []byte(str)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(str) {
		fmt.Printf("%d-%c ", i, ch)
	}
	fmt.Println()

	// 字符串的 rune 数量
	fmt.Println(utf8.RuneCountInString(str))

	// 总结
	// 使用 len 可以获得 字节长度
	// 使用 []byte() 可以获得所有字节组成数组的一个 byte slice
	// 使用 []rune() 会将所有字节转换好后放到一个数组，然后再返回一个 rune slice

	// 字符串的其他操作，看 strings 这个包
	fmt.Println(strings.Index("123", "30"))

	fmt.Println("----------------------")

	// Fields, Split, Join
	fmt.Println(strings.Fields("i love             you"))
	fmt.Println(strings.Split("abcdeabcd", "c"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	// Contains, Index
	fmt.Println(strings.Contains("abc", "cd"))
	fmt.Println(strings.Index("abc", "c"))
	// ToLower, ToUpper
	fmt.Println(strings.ToLower("AbC"))
	fmt.Println(strings.ToUpper("AbC"))
	// Trim, TrimRight, TrimLeft
	fmt.Println(strings.Trim(" 12 34 ", " "), "--")
	fmt.Println(strings.TrimRight(" 12 34 ", " "), "--")
	fmt.Println(strings.TrimLeft(" 12 34 ", " "), "--")
}
