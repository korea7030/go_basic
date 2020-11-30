package main

import "fmt"

func main() {
	for dan := 1; dan <= 9; dan++ {
		fmt.Printf("%d 단 \n", dan)

		for num := 1; num <= 9; num++ {
			if dan == 3 && num == 2 {
				continue
			} else {
				fmt.Printf("%d * %d = %d\n", dan, num, dan*num)
			}
		}
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 3; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 5-2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
