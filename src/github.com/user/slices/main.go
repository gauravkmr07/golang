package main

import "fmt"

func main() {
	fmt.Println("Slices Examples")
	s := make([]string, 3)
	// fmt.Println(s)

	s[0] = "A"
	s[1] = "B"
	s[2] = "C"

	// fmt.Println(s)
	// fmt.Println(len(s))

	s = append(s, "D")
	// fmt.Println(s)
	// fmt.Println(len(s))

	s = append(s, "E", "F", "G")
	// fmt.Println(s)
	// fmt.Println(len(s))

	c := make([]string, len(s))

	copy(c, s)

	c = append(c, "H")
	// fmt.Println(s)
	// fmt.Println(len(s))
	fmt.Println(c)
	//fmt.Println(len(c))

	l := c[0:1]
	fmt.Println(l)
}
