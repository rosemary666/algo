package main

import (
	"fmt"
	"sort"
)

var (
	res  = make([][]int, 0) // 最终结果集
	path = make([]int, 0)   // 中间临时结果
)

func backtracking(candidates []int, target int, startIndex int, sum int) {
	// 超过目标值, 回溯终止
	if sum > target {
		return
	}
	// 找到目标值, 将临时结果追加到最终结果中
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	// for i := startIndex; i < len(candidates); i++ { // 未剪枝
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		backtracking(candidates, target, i, sum) // 从i而不是i+1, 表示可以继续重复
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	sort.Ints(candidates) // 数组先排序, 用于剪枝优化
	backtracking(candidates, target, 0, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", combinationSum([]int{2, 3, 6, 7}, 7))

	println("UseCase 2......")
	fmt.Printf("%+v\n", combinationSum([]int{2, 3, 5}, 8))

	println("UseCase 3......")
	fmt.Printf("%+v\n", combinationSum([]int{2}, 1))
}
