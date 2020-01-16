package main

import "fmt"

func main() {
	fmt.Println("Pointer Examples")
	num := 10
	fmt.Println("Call By Value ", callByVal(num))
	fmt.Println("Call By Ref ", callByRef(&num))
	fmt.Println("Call By Ref ", callByRef(&num))
	fmt.Println("Call By Ref ", callByRef(&num))
}

func callByRef(num *int) int {
	return *num + 1
}

func callByVal(num int) int {
	return num + 1
}
