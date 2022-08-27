package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := []int{}
	m := make(map[int]int)
	for index, v := range nums {
		preIndex, ok := m[target-v]
		if ok {
			res = append(res, preIndex, index)
			break
		}
		m[v] = index
	}
	return res
}

func main() {
	println("UseCase 1......")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("%v\n", twoSum(nums, target))

	println("UseCase 2......")
	nums = []int{3, 2, 4}
	target = 6
	fmt.Printf("%v\n", twoSum(nums, target))

	println("UseCase 3......")
	nums = []int{3, 3}
	target = 6
	fmt.Printf("%v\n", twoSum(nums, target))
}
