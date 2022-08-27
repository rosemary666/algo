package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{} // 用数组实现一个栈

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			// 弹出栈头两个元素
			if len(stack) < 2 {
				return 0
			}
			firstNum := stack[len(stack)-1]
			secondNum := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			ret := 0
			switch token {
			case "+":
				ret = secondNum + firstNum
			case "-":
				ret = secondNum - firstNum
			case "*":
				ret = secondNum * firstNum
			case "/":
				ret = secondNum / firstNum
			}
			stack = append(stack, ret)
		}
	}
	return stack[0]
}

func main() {
	println("UseCase 1......")
	token := []string{"2", "1", "+", "3", "*"}
	println(evalRPN(token))

	println("UseCase 2......")
	token = []string{"4", "13", "5", "/", "+"}
	println(evalRPN(token))

	println("UseCase 3......")
	token = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	println(evalRPN(token))
}
