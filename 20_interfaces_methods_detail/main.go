package main

import (
	"fmt"
	"math"
)

func main() {
	f1()
}

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func f1() {
	p := Point{1, 1}
	q := Point{5, 4}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance // Method Value; value of p is copied just here, no matter if p changes later, value of p copied remains fixed
	fmt.Printf("%T\n", distanceFromP)

	fmt.Println(distanceFromP(q))
	p = Point{2, 2}
	fmt.Println(distanceFromP(q)) // same value, no change
	// ** but if Distance takes a pointer receiver(i.e p *Point) instead of value receiver, then we will see the change in the value,
	// because now the pointer to p is copied not the value

}
