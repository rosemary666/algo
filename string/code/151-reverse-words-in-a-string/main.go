package main

import "fmt"

func reverseWords(s string) string {
	b := []byte(s)
	bLen := len(b)

	/*---------------------1. 移除空格(开头、结尾以及单词之间多余一个的空格)----------------------*/
	slowIndex := 0 // 定义慢指针
	fastIndex := 0 // 定义快指针
	for ; fastIndex < bLen; fastIndex++ {
		// 遇到非空字符, 要考虑单词前面还需要填充一个空格
		if b[fastIndex] != ' ' {
			// 非字符串开头遇到空格(判断fastIndex前一位是不是空格, 是空格即是单词的第一个开头字符)
			if slowIndex != 0 && b[fastIndex-1] == ' ' {
				b[slowIndex] = ' '
				slowIndex++
			}
			b[slowIndex] = b[fastIndex]
			slowIndex++
		}
	}
	// 重新调整数组大小
	b = b[:slowIndex]

	/*---------------------2. 进行整个字符串的反转----------------------*/
	reverseString(b, 0, len(b)-1)
	/*---------------------3. 进行单词内部之间的反转----------------------*/
	start := 0
	end := 0
	for ; end < len(b); end++ {
		// 碰到空字符, 代表从start到end之间是一个单词
		if b[end] == ' ' {
			reverseString(b, start, end-1)
			start = end + 1
		}
		// 代表是最后一个字符
		if end == len(b)-1 {
			reverseString(b, start, end)
		}
	}
	return string(b)
}

func reverseString(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func main() {
	println("UseCase 1......")
	s := "the sky is blue"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reverseWords(s))

	println("UseCase 2......")
	s = "  hello world  "
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reverseWords(s))

	println("UseCase 3......")
	s = "a good   example"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", reverseWords(s))
}
