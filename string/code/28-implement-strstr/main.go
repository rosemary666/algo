package main

import "fmt"

// 获取next数组
func getNextArr(needle string) (next []int) {
	next = make([]int, len(needle))
	cur := 0   //cur是指向当前位置的，代表要求以cur位置字符结尾的最长相同前后缀
	recur := 0 //recur有两个含义，可能是上一轮循环之后传过来的（代表以前一个字符结尾的最长相同前后缀），也可能是递归的时候递归到的，反正也不用管这么多，**反正recur在赋值给next时候就会指向最长相同前缀的下一位**
	next[0] = 0
	for cur = 1; cur < len(needle); cur++ {
		// 如果s[cur]和s[recur]不相等,需要不停的进行回退
		for recur > 0 && needle[cur] != needle[recur] {
			recur = next[recur-1]
		}
		// 如果相等，代表最长公共前后缀+1
		if needle[cur] == needle[recur] {
			recur++
		}
		next[cur] = recur
	}
	return next
}

func strStr(haystack string, needle string) int {
	// 对于本题而言，当 `needle` 是空字符串时我们应当返回 `0` 。这与 C 语言的 `strstr()` 以及 Java 的 `indexOf()` 定义相符。
	if len(needle) == 0 {
		return 0
	}
	j := 0 // 指向next数组
	next := getNextArr(needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	// 不存在返回-1
	return -1
}

func main() {
	println("UseCase 1......")
	haystack := "hello"
	needle := "ll"
	fmt.Printf("%v\n", strStr(haystack, needle))

	println("UseCase 2......")
	haystack = "aaaaa"
	needle = "bba"
	fmt.Printf("%v\n", strStr(haystack, needle))
}
