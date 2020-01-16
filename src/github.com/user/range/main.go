package main

import "fmt"

func main() {
	fmt.Println("Range Examples")

	var marks = [3]int{10, 20, 30}
	fmt.Println(sum(marks))

	var str string = "I am learning golang!"

	charBychar(str)
}

func charBychar(str string) {
	for _, s := range str {
		fmt.Println(string(s))
	}
}

func sum(marks [3]int) int {
	var sum int = 0
	for _, mark := range marks {
		sum = sum + mark
	}

	return sum
}
