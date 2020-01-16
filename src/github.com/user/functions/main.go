package main

import "fmt"

func main() {
	fmt.Println("Functions examples!")
	fmt.Println("Add func called ", add(10, 20.10))
	fmt.Println(multReturn())

	err, integer := multReturn()

	fmt.Println("Error ::", err)
	fmt.Println("Integer ::", integer)

	if err == "" {
		fmt.Println("No error")
	}

	slc := []int{10, 20, 30}
	fmt.Println("---- Variadic Functions ----")
	fmt.Println("### Sum of int values ###", sum(10, 20, 30, 40, 50))
	fmt.Println("### Sum of int values in slice ###", sum(slc...))
}

func sum(nums ...int) int {
	var sum int = 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func multReturn() (string, int) {
	return "", 7
}

func add(a, b float32) float32 {
	return float32(a + b)
}
