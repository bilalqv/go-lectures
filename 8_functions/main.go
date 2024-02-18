package main

import "fmt"

func main() {

	// ----------------- pass Array to function ----------------
	// a := [...]int{1, 2, 3}
	// fmt.Println("a: ", a)
	// b := doArr1(a) // array descriptor is copied. b is completely a different array
	// fmt.Println("b: ", b)
	// fmt.Println("a: ", a) // the original array is not modified

	// ---------------- pass slice to function ----------------
	// a := []int{1, 2, 3}
	// fmt.Println("a: ", a)
	// fmt.Printf("a @ %p\n", a)
	// b := doSlice1(a) // array descriptor is copied. b is completely a different array
	// fmt.Println("b: ", b)
	// fmt.Println("a: ", a) // the original slice gets modified

	// ------------------ pass Map to function ----------------
	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m: ", m)
	do1(m)
	fmt.Println("m: ", m) // at this point m is the same map as m1. The original m created above is gone, we overwrote it.
	// ------------------------
	// mm := map[int]int{4: 1, 7: 2, 8: 3}
	// fmt.Println("mm: ", mm)
	// do2(&mm)
	// fmt.Println("mm: ", mm)

	// -------------------------------- Defer -------------------------

	defer defer1()
	defer defer2()
	fmt.Println("Inside main")

}

// values are copied. b is completely a different array
func doArr1(b [3]int) int {
	b[0] = 9
	fmt.Println("b: ", b)
	return b[0]
}

func doSlice1(b []int) int {
	fmt.Printf("b @ %p\n", b)
	b[0] = 9
	fmt.Println("b: ", b)
	return b[0]
}

// map is passed by reference, map descriptor is copied to m1
func do1(m1 map[int]int) {
	m1[3] = 9              // the original map m is modified. the m1 descriptor's hashmap gets modified which points to original hashmap of m at this point.
	m1 = make(map[int]int) // now we change the m1 descriptor to another hashmap, while original m descriptor is still the same. Now m has no relation with m1
	m1[4] = 7
	fmt.Println("m1: ", m1)
}

// reference(address) of map descriptor is passed
func do2(m1 *map[int]int) {
	(*m1)[3] = 9            // the original map m is modified
	*m1 = make(map[int]int) // we're not only changing the map descriptor in do2, I'm changing the map descriptor in main also
	(*m1)[4] = 7
	fmt.Println("m1: ", *m1)
}

func defer1() {
	fmt.Println("Defer 1 Called")
}
func defer2() {
	fmt.Println("Defer 2 Called")
}
