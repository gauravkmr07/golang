package main

import "fmt"

func main() {
	fmt.Println("For Loop Examples")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}

	var nums = []int{101, 102, 103, 104}

	for j, num := range nums {
		fmt.Println(j, num)
	}

	l := 1

	for {
		fmt.Println("Hello Go!!")
		if l == 4 {
			break
		}

		l++
	}
}
