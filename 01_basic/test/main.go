package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 字符串操作
	str := "Hello, 世界!"
	fmt.Println("String:", str)
	fmt.Println("Length:", len(str)) // 字符串长度

	// 遍历字符串
	for i, c := range str {
		fmt.Printf("Character %c at index %d\n", c, i)
	}

	// 字符操作
	r := '好'
	fmt.Printf("Rune: %c, Unicode code point: %U\n", r, r)

	// Rune 序列长度
	fmt.Println("Rune count:", utf8.RuneCountInString(str))
}
