package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	timerChan := time.After(10 * time.Second)   // 일정 입력시간후에 channel이 나오게 하는것
	tickTimerChan := time.Tick(2 * time.Second) // 일정 시간 주기별로 channel이 나오게 하는 것

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
		}
	}
}
