// run command: go run ./find/ xkcd.json someone bed sleep

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// { "month":      "4",
//   "day":        "20"
//   "year":       "2009",
//   "num":        571,
//   . . .
//   "transcript": "[[Someone is in bed, . . . long int.",
//   "img":        "https://imgs.xkcd.com/comics/cant_sleep.png",
//   "title":      "Can't Sleep",
// }

type xkcd struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	fn := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "no search terms given")
		os.Exit(0)
	}

	// if input, err = os.Open(fn); err != nil {
	// 	fmt.Fprintf(os.Stderr, err)
	// 	os.Exit(-1)
	// }

	var (
		input io.ReadCloser
		items []xkcd
		terms []string
		cnt   int
		err   error
	)

	// some code here
	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, "bad file: %s\n", err)
		os.Exit(-1)
	}

	// decode file
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "bad JSON: %s\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items))

	// get search params
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	if len(terms) < 1 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}
	// search
outer:
	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)
		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}

		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %q\n", item.Num, item.Month, item.Day, item.Year, item.Title)
		cnt++
	}

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)
}
