package main

import (
	"fmt"
	"math"
)

//interfaces in go are a collection of method definitions

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Printf("\n")

}

func main() {
	r := rect{width: 4, height: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
