package main

import "fmt"

type Student struct {
	name  string
	class int

	score Score
}

type Score struct {
	name  string
	grade string
}

// 구조체의 함수
func (s Student) ViewScore() {
	fmt.Println(s.score)
}

func (s Student) InputScore(name string, grade string) {
	s.score.name = name
	s.score.grade = grade
}

func main() {
	var s Student
	s.name = "ljh"
	s.class = 1

	s.score.name = "수학"
	s.score.grade = "C"

	s.ViewScore()
	s.InputScore("과학", "A") // 값이 복사(주소가 아님)
	s.ViewScore()
}
