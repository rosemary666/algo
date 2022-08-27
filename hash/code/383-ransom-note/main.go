package main

func canConstruct(ransomNote string, magazine string) bool {
	// 小写字母总共只有26位，所以定义长度为26的数组大小
	arr := [26]int{}
	// 遍历magazine字符串, ($v-`a`)代表该字符在数组的索引，同时把数组下该索引对应的值加1
	for _, v := range magazine {
		arr[v-rune('a')]++
	}
	// 遍历ransomNote字符串，($v-`a`)代表该字符在数组的索引, 如果该值大于0，则代表可以构成，同时该数组索引对应的值减1
	for _, v := range ransomNote {
		if arr[v-rune('a')] <= 0 {
			return false
		}
		arr[v-rune('a')]--
	}
	return true
}

func main() {
	println("UseCase 1......")
	ransomNote := "a"
	magazine := "b"
	println(canConstruct(ransomNote, magazine))

	println("UseCase 2......")
	ransomNote = "aa"
	magazine = "ab"
	println(canConstruct(ransomNote, magazine))

	println("UseCase 3......")
	ransomNote = "aa"
	magazine = "aab"
	println(canConstruct(ransomNote, magazine))
}
