package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{name, age}
}

func main() {
	fmt.Println("Structs Examples")
	var p person = person{"Ram", 25}
	fmt.Println(p.name)
	fmt.Println(person{"Bob", 20})
	var p1 = newPerson("rajat", 30)
	fmt.Println(p1.name)

	p.name = "Raman"

	fmt.Println(p)
}
