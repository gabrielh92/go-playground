package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func main() {
	// testing code
	tri := triangle{base: 3, height: 5}
	sq := square{side: 7}

	printArea(tri) // exoected 7.5
	printArea(sq)  // expected 49
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.side * s.side
}
