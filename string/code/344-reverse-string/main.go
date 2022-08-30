package main

import "fmt"

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return
}

func main() {
	println("UseCase 1......")
	s := []byte("hello")
	reverseString(s)
	fmt.Printf("%s\n", s)

	println("UseCase 2......")
	s = []byte("Hannah")
	reverseString(s)
	fmt.Printf("%s\n", s)
}
