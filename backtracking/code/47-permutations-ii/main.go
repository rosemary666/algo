package main

import (
	"fmt"
	"sort"
)

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放中间结果
	used = make([]int, 0)   // 判断元素是否已经使用过
)

func backtracking(nums []int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// used[i - 1] == 1，说明同一树枝nums[i - 1]使用过
		// used[i - 1] == 0，说明同一树层nums[i - 1]使用过
		// 如果同一树层nums[i - 1]使用过则直接跳过
		// 这里是去重!!!
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
			continue
		}
		// 同一层使用过
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

func permuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]int, len(nums))
	// ！！！！！！！去重一定要排序
	sort.Ints(nums)
	backtracking(nums)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", permuteUnique([]int{1, 1, 2}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", permuteUnique([]int{1, 2, 3}))
}
