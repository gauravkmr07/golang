package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return r.height*2 + r.width*2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Println("Interface Examples")
	rect := rectangle{width: 10, height: 10}
	circ := circle{radius: 10}
	measure(rect)
	measure(circ)
}
