package main

import "fmt"

func main() {
	fmt.Println("Arrays Example")

	var a [5]int
	fmt.Println(a)

	a[4] = 50
	fmt.Println("Set :: ", a)
	fmt.Println("Get :: ", a[4])

	fmt.Println("Single dimensional Integer array")
	b := [3]int{1, 22, 333}
	fmt.Println(b)

	fmt.Println("Multi dimensional integer array")

	c := [3][3]int{{11, 22, 33}, {44, 55, 66}, {77, 88, 99}}
	fmt.Println(c)

	for _, val1 := range c {
		for _, val2 := range val1 {
			fmt.Print(val2, ", ")
		}

		fmt.Println()
	}

	d := [...]string{"A", "B", "C", "D"}

	fmt.Println(len(d))

}
