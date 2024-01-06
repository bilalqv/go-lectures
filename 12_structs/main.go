package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	// ---------------------- map with struct values ------------
	// f1()

	// ----------------- struct tags --------------------
	// f2()
}

func f1() {
	myMap := map[string]*Employee{}

	myMap["bhat"] = &Employee{"bhat", 2, nil, time.Now()}

	myMap["abc"] = &Employee{
		Name:   "abc",
		Number: 1,
		Boss:   myMap["bhat"],
		Hired:  time.Now(),
	}

	myMap["bhat"].Number++ // I can do it only because map values are pointers, otherwise not.

	fmt.Printf("%T %+[1]v\n", myMap["bhat"])
	fmt.Printf("%T %+[1]v\n", myMap["abc"])
}

func f2() {
	r := Response{Page: 1, Words: []string{"up", "in", "out"}}

	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response
	_ = json.Unmarshal(j, &r2) // we need to pass the pointer where to store i.e &r2
	fmt.Printf("%#v\n", r2)

}
