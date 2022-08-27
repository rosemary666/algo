package main

func isHappy(n int) bool {
	set := make(map[int]struct{})

	for n != 1 {
		n = calc_n2(n)
		_, ok := set[n]
		if ok {
			return false
		}
		set[n] = struct{}{}
	}
	return true
}

// 计算n各个位的平方
func calc_n2(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func main() {
	println("UseCase 1......")
	n := 19
	println(isHappy(n))

	println("UseCase 2......")
	n = 2
	println(isHappy(n))
}
