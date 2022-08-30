package main

func removeDuplicates(s string) string {
	n := len(s)
	stack := []byte{} //用数组模拟栈

	for i := 0; i < n; i++ {
		// 栈非空，并且该字符等于栈头元素, 则栈头元素出栈
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			// 直接入栈
			stack = append(stack, s[i])
		}
	}
	// 这里还是模拟栈实现, 就不直接用原数组了
	// 将栈内剩余元素取出，并倒序排序即可。
	elements := make([]byte, len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		// 弹出栈首元素
		x := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 栈首元素从数组最后开始往前存放
		elements[i] = x
	}
	return string(elements)
}

func main() {
	println("UseCase 1......")
	println(removeDuplicates("abbaca"))
}
