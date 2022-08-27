package main

import "fmt"

func generateMatrix(n int) [][]int {
	var (
		spiral_matrix              = make([][]int, n) //螺旋矩阵
		spiral_matrix_elements_sum = n * n            // 螺旋矩阵元素总数

		left   = 0     // 左边起始位置
		right  = n - 1 // 右边起始位置
		top    = 0     // 上边起始位置
		bottom = n - 1 //下边起始位置

		current_num = 1 // 螺旋矩阵当前记录的元素(初始从1开始)
	)
	for i := 0; i < n; i++ {
		spiral_matrix[i] = make([]int, n)
	}

	for current_num <= spiral_matrix_elements_sum {
		// 从左到右
		for i := left; i <= right; i++ {
			spiral_matrix[top][i] = current_num
			current_num++
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			spiral_matrix[i][right] = current_num
			current_num++
		}
		right--

		// 从右到左
		for i := right; i >= left; i-- {
			spiral_matrix[bottom][i] = current_num
			current_num++
		}
		bottom--

		// 从下到上
		for i := bottom; i >= top; i-- {
			spiral_matrix[i][left] = current_num
			current_num++
		}
		left++
	}
	return spiral_matrix
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
}
