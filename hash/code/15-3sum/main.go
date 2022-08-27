package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 对数组进行从小到大的排序
	sort.Ints(nums)
	res := make([][]int, 0)

	// 定义三个数begin, left, rigth(i, left, right不能相同)
	for i := 0; i < len(nums)-2; i++ { // begin最多到len(nums)-2
		begin := nums[i]
		// 去除begin的重复值
		if i > 0 && begin == nums[i-1] {
			continue
		}
		left := i + 1          // 左指针指向begin前一位
		right := len(nums) - 1 // 右指针指向最后一位
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// left去重
				curLeftValue := nums[left]
				for left < right && curLeftValue == nums[left] {
					left++
				}
				// right去重
				curRightValue := nums[right]
				for left < right && curRightValue == nums[right] {
					right--
				}
			} else if sum > 0 {
				// 因为数组已经排序, 所以right应该向数值小的地方移动
				right--
			} else {
				// 因为数组已经排序, 所以left应该向数值大的地方移动
				left++
			}
		}

	}
	return res
}

func main() {
	println("UseCase 1......")
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v\n", threeSum(nums))

	println("UseCase 2......")
	nums = []int{}
	fmt.Printf("%v\n", threeSum(nums))

	println("UseCase 3......")
	nums = []int{0}
	fmt.Printf("%v\n", threeSum(nums))
}
