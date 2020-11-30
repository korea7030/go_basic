package main

import "fmt"

func main() {
	// 한글은 3byte
	s := "Hello 월드"

	s2 := []rune(s)

	fmt.Println("len(s) = ", len(s))
	fmt.Println("len(s2) = ", len(s2)) // rune 배열(들어오는 내용에 따라 크기가 달라짐)

	for i := 0; i < len(s2); i++ {
		// 문자의 정수값
		fmt.Print(s2[i], ", ")
	}
}
