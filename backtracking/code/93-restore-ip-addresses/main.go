package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	res  = make([]string, 0) // 最终结果, 存储的有效ip地址列表
	path = make([]string, 0) // 存储的中间结果
)

func isValidIp(s string) bool {
	sRune := []rune(s)
	start := 0
	end := len(sRune) - 1
	if start > end {
		return false
	}
	// 首字母为0，但长度大于1
	if end > start && sRune[start] == '0' {
		return false
	}
	sInt, err := strconv.Atoi(string(sRune[start : end+1])) // 注意，这里包含end，所以在golang的切片中需要end+1
	if err != nil {
		return false
	}
	if sInt > 255 {
		return false
	}
	return true
}

func backtracking(s string, startIndex int) {
	sRune := []rune(s)
	// 走到了终点
	if startIndex == len(sRune) && len(path) == 4 {
		// 用.拼接起来
		res = append(res, strings.Join(path, "."))
		return
	}
	for i := startIndex; i < len(sRune); i++ {
		partS := string(sRune[startIndex : i+1]) // 从[startIndex, i]区间的字符串
		// 这里可以进一步剪枝, 子字符串长度不超过3, 当前path不超过4, 并且是有效ip
		if (i-startIndex) <= 2 && len(path) < 4 && isValidIp(partS) {
			path = append(path, partS)
		} else {
			continue
		}
		backtracking(s, i+1)
		path = path[:len(path)-1]
	}
}

func restoreIpAddresses(s string) []string {
	res = make([]string, 0)
	path = make([]string, 0)
	backtracking(s, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", restoreIpAddresses("25525511135"))

	println("UseCase 2......")
	fmt.Printf("%+v\n", restoreIpAddresses("0000"))

	println("UseCase 3......")
	fmt.Printf("%+v\n", restoreIpAddresses("101023"))
}
