package main

import "fmt"

func reverseStr(s string, k int) string {
	b := []byte(s)

	end := len(b)

	for start := 0; start < end; start += 2 * k {
		sub := end - start
		//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样
		if sub >= k {
			reverse_part(b, start, start+k-1)
		} else {
			//如果剩余字符少于 k 个，则将剩余字符全部反转。
			reverse_part(b, start, start+sub-1)
		}
	}
	return string(b)
}

// 反转指定部分的数组元素
func reverse_part(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

func main() {
	println("UseCase 1......")
	s := "abcdefg"
	fmt.Printf("%s\n", reverseStr(s, 2))

	println("UseCase 2......")
	s = "abcd"
	fmt.Printf("%s\n", reverseStr(s, 4))
}
