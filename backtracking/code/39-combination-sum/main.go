package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	return nil
}

func main() {
	println("UseCase 1......")
	fmt.Printf("%+v\n", combinationSum([]int{2, 3, 6, 7}, 7))

	println("UseCase 2......")
	fmt.Printf("%+v\n", combinationSum([]int{2, 3, 5}, 8))

	println("UseCase 3......")
	fmt.Printf("%+v\n", combinationSum([]int{2}, 1))
}
