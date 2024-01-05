package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// --------------------- io functions -------------
	// fmt.Println("print to standard output\n")
	// fmt.Fprintln(os.Stderr, "printing to standard error\n")

	// // always outputs to os.Stdout
	// fmt.Println("Println output\n") // it puts all arguments to one line
	// fmt.Printf("Printf output\n")   // same but here first argument is a formatting string which tells how we want the output to be formatted

	// // these allow us to specify an output identifier to which we want to print to
	// fmt.Fprintln(os.Stderr, "Fprintln\n")
	// fmt.Fprintf(os.Stderr, "Fprintf\n")

	// // these don't print to output but return the formatted string
	// fmt.Sprintln("Sprintln\n")
	// fmt.Sprintf("Sprintln\n")

	// // ---------------------------- format codes -------------
	// a, b := 23, 1435
	// c, d := 12.4, 4.5129

	// fmt.Printf("%d %d \n", a, b)
	// fmt.Printf("%x %x \n", a, b)
	// fmt.Printf("%#x %#x \n", a, b)
	// fmt.Printf("%f %.3f \n", c, d)

	// fmt.Printf("| %7d | %7d |\n", a, b) // format with column size
	// fmt.Printf("| %07d | %07d |\n", a, b)
	// fmt.Printf("| %-7d | %-7d |\n", a, b) // left allign

	// -------------------------- File IO   = linux cat program --------------------------
	// run command: go run . *.txt                  or                go run . *.txt > d.txt
	// for _, fname := range os.Args[1:] {
	// 	file, err := os.Open(fname)

	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		continue
	// 	}

	// 	// Copy takes what is the the 'file' and puts it into the target here 'Stdout'
	// 	if _, err := io.Copy(os.Stdout, file); err != nil {
	// 		fmt.Fprint(os.Stderr, err)
	// 		continue
	// 	}

	// 	file.Close()
	// }

	// --------------------------- program to get the file size --------------------
	calculateSize1()
	wcLinuxProgram()

}

func calculateSize1() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// ReadAll returns a slice of bytes & and error
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")

		file.Close()
	}
}

// linux wc - word count program
// TODO: support arguments like wc , eg: -l etc
func wcLinuxProgram() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		// .Scan calls for one line each time
		for scan.Scan() {
			s := scan.Text() // gives the text of the line
			cc += len(s)
			wc += len(strings.Fields(s)) // .Fields splits the line at spaces and returns []string
			lc++

		}

		fmt.Printf("%7d %7d %7d %7s\n", lc, wc, cc, fname)

		file.Close()
	}
}
