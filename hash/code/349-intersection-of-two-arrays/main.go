package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	arr := make([]int, 0)

	// 数组num1塞入set之中
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	//遍历数组num2, 找到相同元素
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			arr = append(arr, v)
			delete(set, v)
		}
	}
	return arr
}

func main() {
	println("UseCase 1......")
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Printf("%v\n", intersection(num1, num2))

	println("UseCase 2......")
	num1 = []int{4, 9, 5}
	num2 = []int{9, 4, 9, 8, 4}
	fmt.Printf("%v\n", intersection(num1, num2))
}
