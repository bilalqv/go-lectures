package main

import "fmt"

func main() {
	f2()
	println("f2 fix")
	f2fix()
}

// -------------- take struct by value(copy) & modify ----------------
type Widget struct {
	ID    int
	Count int
}

func Expend(w Widget) Widget {
	w.Count-- // change it
	return w  // return it
}

// -------------------------- creating slices from arrays *Important -------------
func f2() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items { // item is a copy of an entry in the items slice i.e copy of the 2 byte array. item is always at a particular location in the memo
		// and everytime in the loop that copy is in the same place
		a = append(a, item[:]) // slice a gets a reference to that fixed location of item, a is picking up the pointer to that location
	} // when the for loop is done, that item location has [5,6] stored, and all entries in a point to the same location having [5,6] stored

	fmt.Println(items)
	fmt.Println(a)

}

func f2fix() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item)) // create new slice each time
		copy(i, item[:])
		a = append(a, i)
	}

	fmt.Println(items)
	fmt.Println(a)

}
