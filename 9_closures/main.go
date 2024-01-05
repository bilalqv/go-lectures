package main

import "fmt"

func main() {

	// // ---------------------- Closure 1 -----------------

	// f := fib()

	// for x := f(); x < 100; x = f() {
	// 	fmt.Println(x)
	// }

	// // -------------------------- Closure 2 --------------
	closure1()
	closure2()
}

// the inner function gets a reference to the outer function's vars
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func closure1() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d  @  %p\n", i, &i)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}
}

func closure2() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // closure capture - simple fix
		s[i] = func() {
			fmt.Printf("%d  @  %p\n", i2, &i2)
		}
	}
	for i := 0; i < 4; i++ {
		s[i]()
	}
}
