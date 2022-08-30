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

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := getNextArr(s)
	// 数组长度减去最长相同前后缀的长度相当于是第一个周期的长度，也就是一个周期的长度，如果这个周期可以被整除，就说明整个数组就是这个周期的循环。
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func main() {
	println("UseCase 1......")
	s := "abab"
	fmt.Printf("%v\n", repeatedSubstringPattern(s))

	println("UseCase 2......")
	s = "aba"
	fmt.Printf("%v\n", repeatedSubstringPattern(s))

	println("UseCase 3......")
	s = "abcabcabcabc"
	fmt.Printf("%v\n", repeatedSubstringPattern(s))
}
