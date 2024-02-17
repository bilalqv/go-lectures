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

	// ---------------- Arrays ------------------
	// var a [3]int
	// a[0] = 1
	// fmt.Println(a, len(a))
	// a[0] = 8
	// fmt.Println(a, len(a))

	// var b [3]int {1,2,3}
	// var c [...]{4,5,6}
}
