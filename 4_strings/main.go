package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ----------------- strings 01 -----------------------

	// s := "✅bilal"

	// fmt.Printf("%8T %[1]v %d\n", s, len(s))
	// fmt.Printf("%8T %[1]v\n", []rune(s))

	// b := []byte(s)
	// fmt.Printf("%8T %[1]v %d\n", b, len(b))

	// ----------------- strings 02 --------------------------
	// s := "the quick  brown fox"

	// a := len(s) // 19
	// fmt.Println("a: ", a)
	// b := s[:3] // "the"
	// fmt.Println("b: ", b)
	// c := s[4:9] // "quick"
	// fmt.Println("c: ", c)
	// d := s[:4] + "slow" + s[9:] // replaces 'quick'
	// fmt.Println("d: ", d)

	// // s[5] = 'a' // syntax error
	// s += "es" // now plural copied
	// fmt.Println("s: ", s)

	// // above, b & c use same storage as s, d is gona point to a completely new memory,
	// // s += "es"; here a block of memory is created with text as s above and append "es" at the end. the earlier value does not go away
	// // becoz we have b,c still pointing to the old memory (original string)

	// ------------ strings 03 ----------------
	// Run: go run main.go this thatt < one.txt
	if len(os.Args) < 3 {
		fmt.Printf("Not enought arguments")
		os.Exit(-1)
	}
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}
