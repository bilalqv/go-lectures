package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type intSlice []int
type myS string

type ByteCounter int // ByteCounter is a writer, because it has a right method (below) like any other io writer. So we can replace any other writer with it.

func main() {
	// ---------------------------  01 -----------------------
	// f1()

	// ----------------------- 02 ---------------------------
	// f2()

	// --------------------  Point, Line, Distance, Distancer -----------------
	// side := Line{Point{1, 3}, Point{4, 6}}
	// plot := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}} // since Path is a slice of points, we don't need to write Point{}

	// PrintDistance(side)
	// PrintDistance(plot)

	// -------------------------- Scaling a Line ---------------------
	side := Line{Point{1, 2}, Point{4, 6}}

	s2 := side.ScaleBy(2)

	fmt.Println(side.Distance())
	fmt.Println(s2.Distance())
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2).Distance()) // we created a Line on the fly and called ScaleBy() on it & then distance
	// becoz ScaleBy() does not need a pointer, and it also returns a Line, & hence we can call Distance() on it.

}

func f1() {
	var v intSlice = []int{1, 2, 3}
	var s myS = "hello"

	fmt.Printf("%T %[1]v\n", s)

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
}

func (is intSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, "#") + "]"
}

func (ms myS) String() string {
	return "|:" + string(ms) + ":|"
}

func f2() {
	var c ByteCounter
	f1, _ := os.Open("a.txt")
	f2 := &c

	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

// --------------------  Point, Line, Distance, Distancer -----------------
// type Point struct {
// 	X, Y float64
// }

// type Line struct {
// 	Begin, End Point
// }

// type Path []Point

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
// }

// func (p Path) Distance() (sum float64) {
// 	for i := 1; i < len(p); i++ {
// 		sum += Line{p[i-1], p[i]}.Distance() // here creating a line on the fly, becoz Distance does not need pointer
// 	}
// 	return sum
// }

// type Distancer interface {
// 	Distance() float64
// }

// func PrintDistance(d Distancer) {
// 	fmt.Println(d.Distance())
// }

// --------------------------------- Scaling a Line ----------------------------
type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func (l Line) ScaleBy(f float64) Line { // create a copy of Line, modify it & return it
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)

	return Line{l.Begin, Point{l.End.X, l.End.Y}}

}
