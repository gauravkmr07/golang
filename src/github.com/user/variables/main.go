package main

import "fmt"

func main() {
	var a = "gaurav"
	fmt.Printf("Variable 'a' of type %T has value %v", a, a)
	fmt.Println()

	var x, y, z int = 11, 22, 33

	fmt.Printf("Variable 'x', 'y', 'z' of type int has values %v, %v, %v ", x, y, z)
	fmt.Println()

	var p int
	var q bool
	var r string

	f := "test"

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(f)
}
