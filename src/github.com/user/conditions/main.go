package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("If Else Examples")

	if 7%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 9; num < 10 {
		fmt.Println("Less than 10")
	}

	fmt.Println("--- In Switch Example 1---")
	switchExample1()
	fmt.Println("--- In Switch Example 2---")
	switchExample2()

}

func switchExample2() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println(time.Saturday)
	default:
		fmt.Println(time.Now().Weekday())
	}
}

func switchExample1() {
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}
}
