package main

import (
	"fmt"
	"sort"
)

var (
	res  = make([][]int, 0) // 存放最终结果
	path = make([]int, 0)   // 存放临时结果
	used = make([]bool, 0)  // 存放是否使用了的标志
)

func backtracking(candidates []int, target int, sum int, startIndex int) {
	if sum > target {
		return
	}
	// 满足条件, 追加到最终结果中
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ { // 剪枝
		// used[i - 1] == true，说明同一树枝candidates[i - 1]使用过
		// used[i - 1] == false，说明同一树层candidates[i - 1]使用过
		// 要对同一树层使用过的元素进行跳过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		used[i] = true
		backtracking(candidates, target, sum, i+1)
		sum -= candidates[i]
		path = path[:len(path)-1]
		used[i] = false
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	// !!!!!先要进行排序
	sort.Ints(candidates)
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(candidates))
	backtracking(candidates, target, 0, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

	println("UseCase 2......")
	fmt.Printf("%+v\n", combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
