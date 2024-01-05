package main

import "fmt"

func main() {
	// ------------------ pass Map to function ----------------
	// m := map[int]int{4: 1, 7: 2, 8: 3}
	// fmt.Println("m: ", m)
	// do1(m)
	// fmt.Println("m: ", m)
	// // ------------------------
	// mm := map[int]int{4: 1, 7: 2, 8: 3}
	// fmt.Println("mm: ", mm)
	// do2(&mm)
	// fmt.Println("mm: ", mm)

	// -------------------------------- Defer -------------------------

	defer defer1()
	defer defer2()
	fmt.Println("Inside main")

}

// map is passed by reference, map descriptor is copied to m1
func do1(m1 map[int]int) {
	m1[3] = 9              // the original map m is modified
	m1 = make(map[int]int) // now we change the m1 ddescriptor to another hashmap, while original m descriptor is still the same. Now m has no relation with m1
	m1[4] = 7
	fmt.Println("m1: ", m1)
}

// reference of map descriptor passed
func do2(m1 *map[int]int) {
	(*m1)[3] = 9            // the original map m is modified
	*m1 = make(map[int]int) // now we change the m1 ddescriptor to another hashmap, while original m descriptor is still the same. Now m has no relation with m1
	(*m1)[4] = 7
	fmt.Println("m1: ", *m1)
}

func defer1() {
	fmt.Println("Defer 1 Called")
}
func defer2() {
	fmt.Println("Defer 2 Called")
}
