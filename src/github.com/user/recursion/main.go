package main

import "fmt"

func main() {
	fmt.Println("Recursion Example")
	f := 5
	fmt.Println("Factorial Of ", f, " is ", fact(f))
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return fact(num-1) * num
}
