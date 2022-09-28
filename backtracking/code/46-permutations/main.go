package main

import "fmt"

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放中间结果集
	used = make([]int, 0)   // 存放数值是否被使用
)

func backtracking(nums []int) {
	// 此时说明找到了一组
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	// 注意每次i都从0开始
	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		path = append(path, nums[i])
		used[i] = 1
		backtracking(nums)
		used[i] = 0
		path = path[:len(path)-1]
	}
}

func permute(nums []int) [][]int {
	res = make([][]int, 0)        // 存放最终结果
	path = make([]int, 0)         // 存放中间结果集
	used = make([]int, len(nums)) // 存放数值是否被使用
	backtracking(nums)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", permute([]int{1, 2, 3}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", permute([]int{0, 1}))

	println("UseCase 3......")
	fmt.Printf("%+v\n", permute([]int{1}))
}
