package main

import (
	"fmt"
	"sort"
)

// 使用快排实现
func topKFrequent(nums []int, k int) []int {
	res := []int{}

	// 使用map统计词频
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	for key, _ := range countMap {
		res = append(res, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(res, func(a, b int) bool {
		return countMap[res[a]] > countMap[res[b]]
	})
	return res[:k]
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%v\n", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	println("UseCase 2......")
	fmt.Printf("%v\n", topKFrequent([]int{1}, 1))
}
