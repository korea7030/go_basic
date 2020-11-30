package main

import "fmt"

func main() {
	a := 21
	c := a % 10
	a = a / 10
	d := a % 10

	fmt.Printf("첫번째 수 : %v 두번째 수 : %v\n", c, d)

	e := 4
	fmt.Println(e << 1)
	fmt.Println(e >> 1)

	var f bool
	f = 3 < 4 && 2 < 5
	fmt.Println(f)
}
