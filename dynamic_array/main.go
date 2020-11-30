package main

import "fmt"

func main() {
	a := make([]int, 0, 8)

	fmt.Printf("len(a) = %d\n", len(a)) //length
	fmt.Printf("cap(a) = %d\n", cap(a)) //capacity

	a = append(a, 123) // capacity가 부족하면 새로운 array를 생성하게 됨

	fmt.Printf("len(a) = %d\n", len(a)) //length
	fmt.Printf("cap(a) = %d\n", cap(a)) //capacity

	// capacity 가 없으면 새로운 array 생성
	b := []int{2, 3}
	c := append(b, 4)

	fmt.Printf("%p %p\n", b, c)

	// 같은 주소를 바라봄
	d := make([]int, 2, 4)
	d[0] = 1
	d[1] = 2
	e := append(d, 3)

	fmt.Printf("d, e %p %p\n", d, e)

	fmt.Println(d)
	fmt.Println(e)

	e[0] = 5
	e[1] = 6

	fmt.Println(d)
	fmt.Println(e)

}
