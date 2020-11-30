package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

// *StructA 의 method
func (a *StructA) AAA(x int) int {
	return x * x
}

// *StructA 의 method
func (a *StructA) BBB(x int) string {
	return "X = " + strconv.Itoa(x)
}

// Struct B //
type StructB struct {
}

// 항상 interface에 선언된 모든 method를 선언해줘야 에러가 안남
func (b *StructB) AAA(x int) int {
	return x * 2
}

func (a *StructB) BBB(x int) string {
	return "X = " + strconv.Itoa(x)
}

// End Struct B //

type StructC struct {
	val string
}

func (c *StructC) String() string {
	return "Val : " + c.val
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	c := &StructC{val: "CCC"}
	fmt.Println(c)
	// var c InterfaceA
	// c = &StructA{}

	// // var d InterfaceA
	// // d = &StructB{}

	// fmt.Println(c.BBB(3))
	// fmt.Println(d.AAA(3))
}
