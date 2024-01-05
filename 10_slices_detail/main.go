package main

import "fmt"

func main() {
	// ----------------------- slices 01 - len, cap ----------------
	// a := [3]int{1, 2, 3}
	// b := a[:1]

	// fmt.Println("a = ", a)
	// fmt.Println("b = ", b)

	// c := b[0:2] // capacity of c = length of the underlying array i.e a
	// // c1 := b[0:4]  // error out of range
	// fmt.Println("c = ", c)

	// d := a[0:1:2] // [i:j:k] -> len j-i, cap k-i
	// fmt.Println("d = ", d)

	// -------------------- slices 02 - ---------------------
	a := [...]int{1, 2, 3}
	b := a[:1]
	c := a[:2]

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 7) // here since len(c) < cap(c) -> we modify the underlying array i.e a also
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

}
