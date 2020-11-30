package main

import "fmt"

// (v1 type, v2 type) return_type { }
func add(x int, y int) int {
	return x + y
}

func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	a := 1
	b := 2
	fmt.Print(add(a, b))

	fmt.Println()
	fmt.Print(fibo(5))
}
