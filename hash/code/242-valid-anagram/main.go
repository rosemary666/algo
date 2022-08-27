package main

func isAnagram(s string, t string) bool {
	var arr [26]int
	//遍历s, 找到字符对应的相应索引, 并对该索引相对的值加1
	for _, v := range s {
		arr[rune(v)-rune('a')] += 1
	}
	//遍历t, 找到字符对应的相应索引, 并对该索引相对的值减1
	for _, v := range t {
		arr[rune(v)-rune('a')] -= 1
	}
	//遍历数组, 如果存在值为非0的, 即代表不是异构
	//两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等
	return arr == [26]int{}
}

func main() {
	println("UseCase 1......")
	s := "anagram"
	t := "nagaram"
	println(isAnagram(s, t))

	println("UseCase 2......")
	s = "rat"
	t = "car"
	println(isAnagram(s, t))
}
