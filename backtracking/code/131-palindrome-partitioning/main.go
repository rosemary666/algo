package main

import "fmt"

// 判断是否是回文
func isPalindrome(s string) bool {
	sRune := []rune(s)
	i := 0
	j := len(sRune) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

var (
	res  = make([][]string, 0) // 存放最终结果
	path = make([]string, 0)   // 存放中间结果
)

func backtracking(s string, startIndex int) {
	sRune := []rune(s)
	// 如果起始位置已经大于s的大小，说明已经找到了一组分割方案了
	if startIndex >= len(sRune) {
		temp := make([]string, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := startIndex; i < len(sRune); i++ {
		partS := sRune[startIndex : i+1] // 从[startIndex, i]这部分区域
		// 是回文
		if isPalindrome(string(partS)) {
			path = append(path, string(partS))
		} else {
			continue
		}
		backtracking(s, i+1)      // 寻找i+1为起始位置的子串
		path = path[:len(path)-1] // 回溯过程，弹出本次已经填在的子串
	}
}

func partition(s string) [][]string {
	res = make([][]string, 0)
	path = make([]string, 0)
	backtracking(s, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", partition("aab"))

	println("UseCase 2......")
	fmt.Printf("%+v\n", partition("a"))
}
