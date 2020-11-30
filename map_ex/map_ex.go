package main

import "fmt"

type Key struct {
	v int
}

type Value struct {
	v int
}

func main() {
	// [key_type][value_type]
	var m map[string]string // 선언
	// var m1 map[int]string
	// var m2 map[Key]Value
	// var m3 map[Key]*Value
	m = make(map[string]string) // 초기화
	m["abc"] = "bbb"

	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])

	fmt.Println("m1[55] = ", m1[55])

	m2 := make(map[int]int)
	m2[4] = 4

	// 값을 설정 안하면 default 값으로 설정됨(0)
	fmt.Println("m2[10] = ", m2[10])

	m2[5] = 0

	fmt.Println("m2[5] = ", m2[5])

	v := m1[53]
	// key값 설정여부 확인 가능
	v1, ok1 := m2[4]
	v2, ok2 := m2[10]

	fmt.Println(v, v1, ok1, v2, ok2)

	delete(m1, 53)
	v, ok3 := m1[53]

	fmt.Println(v, ok3)

	m2[2] = 2
	m2[10] = 10

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}
}
