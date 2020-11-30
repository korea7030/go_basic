package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Println("연산자를 입력")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
		break
	case "-":
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
		break
	case "*":
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
		break
	case "/":
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
		break
	default:
		fmt.Println("잘못된 입력")
	}

}
