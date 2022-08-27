package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := map[int]int{}
	// 遍历num1和num2, 统计两个数组元素之和，和出现的次数, 并且存入map之中
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	count := 0
	// 遍历num3和num4
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if times, ok := m[0-(v3+v4)]; ok {
				count += times
			}
		}
	}
	return count
}

func main() {
	println("UseCase 1......")
	num1 := []int{1, 2}
	num2 := []int{-2, -1}
	num3 := []int{-1, 2}
	num4 := []int{0, 2}
	fmt.Printf("%v\n", fourSumCount(num1, num2, num3, num4))

	println("UseCase 2......")
	num1 = []int{0}
	num2 = []int{0}
	num3 = []int{0}
	num4 = []int{0}
	fmt.Printf("%v\n", fourSumCount(num1, num2, num3, num4))
}
