package main

import "fmt"

func main() {
	fmt.Println("Map Data Structure Examples")

	m := make(map[string]int)

	m["k1"] = 101
	m["k2"] = 102
	m["k3"] = 103
	m["k4"] = 104
	m["k5"] = 105
	m["k6"] = 106

	// printVal(m)

	delete(m, "k2")
	printVal(m) // will print in random order

	val, prs := m["K201"]

	if prs {
		fmt.Println("value : ", val)
	} else {
		fmt.Println("Not present!!")
	}

	m2 := map[string]int{"a": 95, "b": 96}

	fmt.Println(m2)

}

func printVal(m map[string]int) {
	for k, val := range m {
		fmt.Println(k, val)
	}
}
