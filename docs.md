## Go Commands
 - go run .

 these two are same commands
 - go run ./3_basic/ < ./3_basic/nums.txt
 - cat ./3_basic/nums.txt | go run ./3_basic/






### Arrays
// slice -> variable length array
// array, slice like strings, are a squence of things, literally laid down in memory one after the another
// map -> elements in map are not necessarily laid down consecutively in memory

All these are equivalent
var a [3]int
var b [3]int{1,2,3}
var c [...]{1,2,3}

var d [3]int
d = b // elements are copied. this is not like strings. If arrays are very large, this is not efficient. That's why Go has another type called slices.

var m [...]int{1,2,3,4,5}
c = m // Type mismatch, becoz these are of different types.

### Slices
a Slice has a descriptor and pointer to some other memory location where the actual data is stored. a slice always has an array behind it.
a slice descriptor has a pointer, length and capacity. 

var a []int
var b = []int{1,2}

a = append(a,1)
append actually copies the elements of the underlying array to a new memory location and then appends the new element at the end, and finally updates the descriptor of a to point to the new memory location.
This is similar to strings where we append some letters to a string.
a = b // overwrites a with b. b descriptor gets copied to the a descriptor. a and b now point to the same underlying array. 

var w = [...]int{1,2,3} // array of length 3
var x = []int{1,2,3} // slice of length 3


### Map
- map is actually a hash table.
- the map's variable name is a descriptor which has a pointer to the underlying hash table.
- assigning one map to another map, copies the descriptor.
- maps are passed by reference.
- maps can't be compared to one another. maps can be compared only to nil.
- nil is a type of zero. it indicates the absence of something.

### Control Statements
- in for loop like: for i,v := range arr.... ; here index and the array elements are copied into i & v respectively. this is not fine if array elements are of larger size
- same is the case with maps
- maps in Go have no order, because it is based on a hash table. It is an un-ordered map.
- inside nested loops, if I want to continue outer loop from inner loop, I can use some label for outer loop and use continue <label> to continue the outer loop from inner loop.
- in switch statement, the default is always looked at last, no matter where we put it in the switch.

### Packages
- every name that is capitalized is exported.
- every source file in the package must import what it needs
- a package "A" can not import a package that imports "A". This will cause a cycle. Go has a Tree structure dependency.
- a package should embed deep functionality behind a simple API.

### Typing
- Go uses structural typing.
- So when do we use named typing? -> we use named typing when we explicitly introduce a type with the type keyword.

### Standard IO
- we have: standard input, output and error

### File IO
- always check for errors and handle them. Don't ignore them.


