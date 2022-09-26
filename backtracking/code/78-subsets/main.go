package main

import "fmt"

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放中间结果
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
		path = append(path, nums[i])
		backtracking(nums, i+1)
		path = path[:len(path)-1]
	}
}

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	backtracking(nums, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", subsets([]int{1, 2, 3}))

	println("UseCase 2......")
	fmt.Printf("%+v\n", subsets([]int{0}))
}
