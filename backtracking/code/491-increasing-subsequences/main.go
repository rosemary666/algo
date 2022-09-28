package main

import "fmt"

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放中间临时结果
)

func backtracking(nums []int, startIndex int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}
	set := make(map[int]int, 0) // 记录该层数值是否已经被使用过
	for i := startIndex; i < len(nums); i++ {
		//分两种情况判断：
		//一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		//或者二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
		if (len(path) > 0 && nums[i] < path[len(path)-1]) || (set[nums[i]] == 1) {
			continue
		}
		path = append(path, nums[i])
		set[nums[i]] = 1
		backtracking(nums, i+1)
		path = path[:len(path)-1]
	}
}

func findSubsequences(nums []int) [][]int {
	res = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)  // 存放中间临时结果
	backtracking(nums, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", findSubsequences([]int{4, 6, 7, 7}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", findSubsequences([]int{4, 4, 3, 2, 1}))
}
