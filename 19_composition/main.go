package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// --------------------------- example 01 ---------------------------
	// f1()

	// --------------------------- Example 02 ---------------------------
	f2()

}

// --------------------------- example 01 ----------------
// type Pair struct {
// 	Path string
// 	Hash string
// }

// type PairWithLength struct {
// 	Pair
// 	Length int
// }

// func f1() {
// 	p1 := PairWithLength{Pair{"/home", "0x12dec2"}, 25}

// 	fmt.Println(p1.Path, p1.Length, p1.Hash) // NOT p1.Pair.Path -> we directly use p1.Path instead of p1.Pair.Path
// 	// because fields of Pair are promoted to the same level as PairWithLength
// }

// --------------------------- Example 02 ---------------------------
type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s, lenght: %d", p.Path, p.Hash, p.length)
}

func Filename(p Pair) string { // it can not take dericed struct as parameter
	return filepath.Base(p.Path)
}

func (p Pair) Filename() string { // method on Pair struct
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func f2() {
	p := Pair{"/usr/home", "0x123wq2"}
	pl := PairWithLength{Pair{"/usr/home", "0x123wq2"}, 25}

	var fn Filenamer = PairWithLength{Pair{"/usr/bin", "0x12sd45"}, 32} // this works becoz PairWithLength has a method Filename() promoted from Pair

	fmt.Println(p)
	fmt.Println(pl)
	fmt.Println(Filename(p)) // I can not pass pl here
	// fmt.Println(Filename(pl)) // Error

	fmt.Println(fn.Filename())

}
