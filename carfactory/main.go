package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}

	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
	}

}

func StarCartWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	// chan type : reference type
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	// fmt.Printf("chan1 : %p\n", &chan1)
	// fmt.Printf("chan2 : %p\n", &chan2)
	// fmt.Printf("chan3 : %p\n", &chan3)

	// chan1 생성
	go StarCartWork(carChan1)
	go StartPlaneWork(planeChan1)

	// chan1 생성 가지고 chan2 생성
	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)

	// chan2 생성 가지고 chan3 생성
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}
}
