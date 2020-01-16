package main

import "fmt"

type emp struct {
	id   int
	name string
	age  int
	dep  string
}

func (e emp) salary() int {
	if e.dep == "admin" {
		return e.age * 2 * 2400
	}
	return e.age * 3 * 2400
}

type rect struct {
	height int
	widht  int
}

func (r rect) area() int {
	return r.widht * r.height
}

func (r rect) perm() int {
	return r.widht*2 + r.height*2
}
func main() {
	fmt.Println("--Methods In Go--")

	r := rect{widht: 10, height: 10}
	fmt.Println(r.area())
	fmt.Println(r.perm())

	rp := &r
	rp.height = 90
	rp.widht = 20

	fmt.Println(rp.area())
	fmt.Println(rp.perm())

	fmt.Println("---- Employee Code ----")

	e := emp{id: 101, name: "gaurav", age: 10, dep: "admin"}
	e2 := emp{id: 102, name: "anoop", age: 30, dep: "tech"}

	fmt.Println(e.salary())
	fmt.Println(e2.salary())
}
