package main

import "fmt"

var (
	path = make([]int, 0)   // 符合条件的结果
	res  = make([][]int, 0) // 存放最终的结果集
)

func backtracking(k int, targetNum int, startIndex int, sum int) {
	if sum > targetNum { // 剪枝操作
		return
	}
	if len(path) == k {
		if sum == targetNum {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	for i := startIndex; i <= 9-(k-len(path))+1; i++ { //裁枝
		path = append(path, i) // 处理
		sum += i               // 处理
		backtracking(k, targetNum, i+1, sum)
		sum -= i                  //回溯
		path = path[:len(path)-1] //回溯
	}
}

func combinationSum3(k int, n int) [][]int {
	path = make([]int, 0)  // 符合条件的结果
	res = make([][]int, 0) // 存放最终的结果集
	backtracking(k, n, 1, 0)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", combinationSum3(3, 7))

	println("UseCase 2......")
	fmt.Printf("%+v\n", combinationSum3(3, 9))

	println("UseCase 3......")
	fmt.Printf("%+v\n", combinationSum3(4, 1))
}
