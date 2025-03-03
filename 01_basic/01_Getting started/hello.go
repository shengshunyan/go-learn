package main

import (
	"fmt"
	// "rsc.io/quote"
)

// func add(a, b int) int {
// 	return a + b
// }

// var c = 30

// d:= 40

func main() {
	// var a = 10
	// var b = a

	// fmt.Println("Hello, World!", a, b, &a, &b)

	// arr := [3]int{1, 2}
	// fmt.Println("arr", arr, arr[2])

	// fmt.Println(quote.Go(), a, add(1, 2))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	siteMap := make(map[string]string)
	siteMap["google"] = "https://www.google.com"
	siteMap["baidu"] = "https://www.baidu.com"

	fmt.Println(siteMap)

}
