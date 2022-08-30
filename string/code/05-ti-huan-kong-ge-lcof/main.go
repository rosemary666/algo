package main

import "fmt"

// 双指针法
func replaceSpace(s string) string {
	b := []byte(s)
	bLen := len(b)
	// 统计空格数
	spaceCount := 0
	for _, v := range b {
		if v == byte(' ') {
			spaceCount += 1
		}
	}
	// 扩容数组
	scalaArr := make([]byte, spaceCount*2) // 这里记得需要乘2
	b = append(b, scalaArr...)

	// 遍历数组进行替换和移动
	// 双指针法, 一个指针指向新数组的末尾，一个指针指向旧指针的末尾
	left := bLen - 1
	right := len(b) - 1
	for left >= 0 {
		// 遇到空格
		if b[left] == ' ' {
			// 填充%20
			b[right] = '0'
			b[right-1] = '2'
			b[right-2] = '%'
			left -= 1
			right -= 3
		} else {
			// 直接替换left索引对应的数组值
			b[right] = b[left]
			left -= 1
			right -= 1
		}
	}
	return string(b)
}

func main() {
	println("UseCase 1......")
	s := "We are happy."
	fmt.Printf("%s\n", replaceSpace(s))
}
