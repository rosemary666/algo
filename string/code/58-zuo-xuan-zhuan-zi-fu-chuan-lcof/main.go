package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	/*-------------1. 反转前n个字符----------------*/
	reverseStr(b, 0, n-1)
	/*-------------2. 反转第n个后面的字符----------------*/
	reverseStr(b, n, len(b)-1)
	/*-------------3. 整体字符串反转----------------*/
	reverseStr(b, 0, len(b)-1)
	return string(b)
}

func reverseStr(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func main() {
	println("UseCase 1......")
	s := "abcdefg"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reverseLeftWords(s, 2))

	println("UseCase 2......")
	s = "lrloseumgh"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reverseLeftWords(s, 6))
}
