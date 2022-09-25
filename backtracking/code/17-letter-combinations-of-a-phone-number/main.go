package main

import "fmt"

var (
	letterMap = [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	s   = make([]rune, 0)   //记录临时结果
	res = make([]string, 0) //记录全局结果
)

func backtracking(digits string, index int) { //index代表遍历到digits的第几个字母，也代表深度
	if index == len(digits) {
		res = append(res, string(s))
		return
	}
	digit := digits[index] - '0' // 将index指向的数字转为int, 比如digits为"23", index为0的情况，去代表取出2对应的字符
	letters := letterMap[digit]  // 取数字对应的字符集
	for i := 0; i < len(letters); i++ {
		s = append(s, rune(letters[i]))
		backtracking(digits, index+1)
		s = s[:len(s)-1]
	}
}

func letterCombinations(digits string) []string {
	s = make([]rune, 0)
	res = make([]string, 0)
	if digits == "" {
		return res
	}
	backtracking(digits, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", letterCombinations("23"))

	println("UseCase 2......")
	fmt.Printf("%+v\n", letterCombinations(""))

	println("UseCase 3......")
	fmt.Printf("%+v\n", letterCombinations("2"))
}
