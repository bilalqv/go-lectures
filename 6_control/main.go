package main

func main() {
	// // --------------------- if else ------------------
	// if res := ifFunc(); res == true {
	// 	fmt.Println("inside if block")
	// }

	// // ------------------------ for loop -----------------
	// myArr := [...]int{11, 22, 33, 44}
	// // here index and the array elements are copied into i & v respectively. this is not fine if array elements are of larger size; so approach 2 is efficiednt
	// for i, v := range myArr {
	// 	fmt.Println(i, v)
	// }

	// // this one is more efficient; arr elements are not copied
	// for i := range myArr {
	// 	fmt.Println(i, myArr[i])
	// }

	// // --------------- loop maps ---------------------
	// myMap := make(map[string]int)
	// myMap["hello"] = 3
	// for k := range myMap { // same as arrats, values are not copied -> Efficient
	// 	fmt.Println(k, myMap[k])
	// }
	// for k, v := range myMap {
	// 	fmt.Println(k, v)
	// }

	// --------------- variable declarations -------------------
	// there are several ways
	// var a int
	// var b int = 1
	// var c = 1
	// var d = 1.2
	// // declaration block
	// var (
	// 	x, y int
	// 	z    string
	// )

	// fmt.Println(a, b, c, d, x, y, z)

	// --------------------- typing ---------------
	// type x int // x is defined type, base int
	// var a x
	// b := 3
	// a = b // type mismatch

}

func ifFunc() bool {
	return true
}
