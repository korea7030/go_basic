package main

import "fmt"

////////////// interface ////////////////
type SpoonOfJam interface {
	String() string
}

// interface
type Jam interface {
	GetOneSpoon() SpoonOfJam
}

////////////// End interface ////////////////

type Bread struct {
	val string
}

/////// Strawberry ///////
type StrawberryJam struct {
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

// strawberryjam 객체의 GetOneSpoon 함수를 SpoonOfJam interface의 method로 구현
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

/////// End Strawberry ///////

////// Orange ///////
type OrangeJam struct {
}

type SpoonOfOrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

////// End Orange ///////

////// Apple //////
type AppleJam struct {
}

type SpoonOfAppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

////// End Apple //////

// bread method
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
