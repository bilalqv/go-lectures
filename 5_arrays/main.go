package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// var a []int
	// var b = []int{1, 2}

	// a = append(a, 1)
	// b = append(b, 3)

	// fmt.Println(a, b, cap(a), cap(b))
	// a = b
	// fmt.Println(a, b, cap(a), cap(b))

	// b = append(b, 8)
	// b = append(b, 9)

	// fmt.Println(a, b, cap(a), cap(b))

	// b[0] = 9
	// fmt.Println(a, b) // we changed underlying array, hence both a & b are updated

	// --------------- function arguments --------------
	// var w = [...]int{1, 2, 3}
	// var x = []int{0, 0, 0}

	// y := do(w, x)
	// fmt.Println(w, x, y)

	// // -------------------- Map -------------------------
	// var m map[string]int      // nil, no storage
	// p := make(map[string]int) // non-nil but empty

	// a := p["the"] // returns 0, we can read and write to a non-nil vlearempty map
	// b := m["the"] // returns 0, we can read but we can not write to a nil map

	// fmt.Println(a, b)

	// // m["and"] = 1  // panic, we can not write to a nil map
	// m = p         // descriptor of p is copied to m, both now point to same map, changing one reflects values in both
	// m["and"]++    // OK, m & p point to same location/table
	// c := p["and"] // returns 1
	// fmt.Println(c)
	// fmt.Printf("map: %v , len: %v\n", p, len(p))

	// // maps have a special two result lookup function, to check if the key was present in the map or not.
	// val1 := p["the"] // returns 0, but is this key's value 0 or the key is not present ??
	// fmt.Println(val1)
	// val, ok := p["the"] // returns 	0, false
	// fmt.Println(val, ok)

	// if val2, ok := p["hello"]; ok {
	// 	// we are sure val2 is not the default value
	// 	fmt.Println(val2)
	// }

	// -------------------- Get top three frequent words in a txt file ---------------
	topFreqWords() // run this with command: go run . < ./freqWords.txt

}

// slices are passed by reference to the function parameters. If we modify the passed slice argument inside the function
// the actual value outside is also changed.
// arrays are passed by value. we create a copy and assign the function parameter to it. original array is not changed
// func do(a [3]int, b []int) []int {
// 	// a = b    // syntax error
// 	a[0] = 4 // w remains unchanged
// 	b[0] = 3 // x changed

// 	c := make([]int, 5)
// 	c[4] = 9
// 	copy(c, b)
// 	return c
// }

func topFreqWords() {
	scan := bufio.NewScanner(os.Stdin)
	wordsMap := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		wordsMap[scan.Text()]++
	}

	fmt.Printf("Total unique words: %v\n", len(wordsMap))

	type kv struct {
		key string
		val int
	}

	var sSlice []kv

	for k, v := range wordsMap {
		sSlice = append(sSlice, kv{k, v})
	}

	sort.Slice(sSlice, func(i, j int) bool {
		return sSlice[i].val > sSlice[j].val
	})

	for _, s := range sSlice[:3] {
		fmt.Println(s.key, " appears ", s.val, " times.")
	}

}
