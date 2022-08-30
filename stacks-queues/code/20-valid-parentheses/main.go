package main

func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	stack := []byte{} // 用数组实现一个栈
	// 字符括号映射表
	charMap := map[byte]byte{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for i := 0; i < n; i++ {
		// 如果是左括号，则直接将其对应的右括号入栈
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, charMap[s[i]])
		} else if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			// 当前字符映射匹配到栈首元素, 栈弹出
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	// 最终判断栈是否为空
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	println("UseCase 1......")
	println(isValid("()"))

	println("UseCase 2......")
	println(isValid("()[]{}"))

	println("UseCase 3......")
	println(isValid("(]"))

	println("UseCase 4......")
	println(isValid("([)]"))

	println("UseCase 5......")
	println(isValid("{[]}"))
}
