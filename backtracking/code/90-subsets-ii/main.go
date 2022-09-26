package main

import (
	"fmt"
	"sort"
)

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放中间结果
	used = make([]bool, 0)  // 记录是否被使用
)

func backtracking(nums []int, startIndex int) {
	// 收集子集，要放在终止添加的上面，否则会漏掉自己
	temp := make([]int, len(path))
	copy(temp, path)
	res = append(res, temp)

	if startIndex >= len(nums) { // 可以不加, for循环会自动退出
		return
	}
	for i := startIndex; i < len(nums); i++ {
		// used[i - 1] == true，说明同一树枝nums[i - 1]使用过
		// used[i - 1] == false，说明同一树层nums[i - 1]使用过
		// 而我们要对同一树层使用过的元素进行跳过
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtracking(nums, i+1)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(nums))
	// 去重需要先排序
	sort.Ints(nums)
	backtracking(nums, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", subsetsWithDup([]int{1, 2, 2}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", subsetsWithDup([]int{0}))
}
