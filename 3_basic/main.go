package main

import (
	"fmt"
	"os"
)

func main() {

	// ------------------- numbers -----------------
	// a := 2
	// b := 3.2

	// fmt.Printf("a: %8T %[1]v\n", a)
	// fmt.Printf("b: %8T %[1]v\n", b)

	// a = int(b)
	// fmt.Printf("a: %8T %[1]v\n", a)

	// b = float64(a)
	// fmt.Printf("b: %8T %[1]v\n", b)

	// -------------  calculate average -------------
	// To read from file run: go run main.go < nums.txt
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++

	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is: ", sum/float64(n))

	// ----------------- strings 01 -----------------------

	// s := "✅bilal"

	// fmt.Printf("%8T %[1]v %d\n", s, len(s))
	// fmt.Printf("%8T %[1]v\n", []rune(s))

	// b := []byte(s)
	// fmt.Printf("%8T %[1]v %d\n", b, len(b))

	// ----------------- strings 02 --------------------------
	// s := "the quick brown fox"

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
	// if len(os.Args) < 3 {
	// 	fmt.Printf("Not enought arguments")
	// 	os.Exit(-1)
	// }
	// old, new := os.Args[1], os.Args[2]
	// scan := bufio.NewScanner(os.Stdin)

	// for scan.Scan() {
	// 	s := strings.Split(scan.Text(), old)
	// 	t := strings.Join(s, new)

	// 	fmt.Println(t)
	// }

	// ---------------- Arrays ------------------
	// var a [3]int
	// a[0] = 1
	// fmt.Println(a, len(a))
	// a[0] = 8
	// fmt.Println(a, len(a))

	// var b [3]int {1,2,3}
	// var c [...]{4,5,6}
}
