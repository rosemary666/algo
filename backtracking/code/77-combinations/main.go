package main

import "fmt"

var (
	path = make([]int, 0)   // 迭代的路径记录
	res  = make([][]int, 0) // 最终结果记录
)

func backtracking(n int, k int, startIndex int) {
	if len(path) == k {
		// 这里拷贝一定一定要当心!!!!!!
		temp := make([]int, k)
		copy(temp, path)
		res = append(res, temp)
		return
	}
	// for i := startIndex; i <= n; i++ { // 未裁枝
	for i := startIndex; i <= n-(k-len(path))+1; i++ { // 控制树的横向遍历
		path = append(path, i)    // 处理节点
		backtracking(n, k, i+1)   // 递归：控制树的纵向遍历，注意下一层搜索要从i+1开始
		path = path[:len(path)-1] // 回溯，撤销处理的节点
	}
}

func combine(n int, k int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)
	backtracking(n, k, 1)
	return res
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", combine(4, 2))

	println("UseCase 2......")
	fmt.Printf("%+v\n", combine(1, 1))
}
