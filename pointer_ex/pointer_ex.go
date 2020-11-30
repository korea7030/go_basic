package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

// 주소값을 통해 입력값이 변경되도록
func (s *Student) PrintScore() {
	fmt.Println(s.class, s.grade)
}

// 주소값을 통해 입력값이 변경되도록
func (s *Student) InputScore(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "aaa", age: 23, grade: "수학", class: "A"}
	s.InputScore("과학", "C")
	s.PrintScore()
}

func Increase(x *int) {
	*x = *x + 1
}
