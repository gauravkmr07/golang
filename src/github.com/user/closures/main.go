package main

import "fmt"

func main() {
	fmt.Println("...Closure Examples...")

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
