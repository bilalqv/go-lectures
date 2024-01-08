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

### Functions
- we can do anything with functions as we can do with string or int
- function signature -> order and type of parameters and return types
- a function declaration lists Formal Parameters, while a function call has Actual Parameters
- Parameters by value -> numbers, bools, arrays, structs
- Parameters by reference -> slices, maps, channels, strings, things passed by pointers (&x)
- since arrays are passed by value, so the actual parameter is copied into the formal parameter. if we modify the arr[i] inside function, the actual parameter will not be modified. 
- Slices are passed by reference, so if we modify the slice inside function, the actual parameter will be modified.
- From a technical point of view, there is no such thing as by reference in Go. Every parameter is passed by value, because the formal parameter is a local variable in the function it gets copied  the value of the actual parameter. NOW, if that value is a descriptor, then we refer to something that was referred to from outside from the function. So, the slice descriptor gets copied to the slice descriptor of the formal parameter. they both point to the actual table of data in memo, so we think of slices by reference, but slice descriptor isn't by reference, it is copied. Similaraly, map descriptor is copied, but they point to the same hash table. Then we have the third case where we pass a pointer to the map descriptor and wehn I change that in function, I am also changing the map descriptor outside the function, so we have our original outside map changed, not just the value inside the map, but it became an entirely different map.
- Go allows returning multiple values from a function.

### Defer
- two consecutive defer statements are executed in reverse order. They get stacked up.
- defer operates at function scope, not block scope. The defer will happen when the function returns, not when the block ends.
- values passed to defer functions get copied at the point of writing the defer statement, not when  defer is executed.

### Closures
- Go approach is to allocate as much as possible on the stack, but it will allocate on the heap when the lifetime will exceed the scope whatever context that variable is in.
- So lifetime can exceed the scope in which  the variable is declared.
- Closure capture.
- a closure is a type of function call that has additional data from outside the function like from another function scope and is closed over by reference.
- Gotcha! -> if a closure if executed asychronously i.e later on then it is possible that the variable that I closed over mutates in the meantime. A simple fix is to create a local copy & then you close over that local copy which can not mutate again

### Slices in Detail
- nil slice -> len = 0, cap = 0, pointer = null
- empty slice -> len = 0, cap = 0, pointer = not null , pointing to a special location in memo
- when we encode a nil slice in json, we get null, but when we encode an empty slice in json, we get an empty array. -> Similaraly with Maps as well.
- to check the slice, we can use: if len(s) == 0 {....} 
- we can append to a nil slice
- a := []int{1, 2, 3} -> here slice a points to some un-named array in memo.


### Structs
- There is a limitation on maps, we can not take the address of a Map entry. When we give a key, we get the value, we can not get the address of the value in the map. The reason being that the entries in the map can re-arrange themselves.
- If I've a map of structs & I want to do something to a value inside that struct like m["key"].field = 5, I can't do that. So, we can use map of say string -> struct pointer, and that will work.
- two named structs with similar fields are not compatible, we can not assign one to another, we first need to type convert them before assignment. But two anonymous structs with similar fields are compatible.
- struct tags are mostly used for doing conversion of data in the program to an external format like json, xml, etc.
- Private fields(lowercse fields) of a struct that are not exported, are not encoded using struct tags.

### Reference & Value
- Pointer Semantics -> Efficient
- Value Semantics -> integrity, especially in concurrency.
- we can not copy mutexes, any struct that contains a mutex, needs to be operated using pointer semantics.
- An alternative to taking a pointer to a struct, is to take the value of the struct, change it & give it back. (This is OK if the struct size is less than 64bytes ) -> common in functional programming.
- generally, stack allocation is more efficient, & Go will try to allocate on stack as much as possible. -> reasons: Garbage collection, how cache works.
- avoid using new in Go.
- Single biggest in Go is anytime u take a reference to a mutating loop variable & u keep it

### Networking
- 

### OOPS
- For most people, essential emelents of OOPS are:
    - Abstraction
    - Encapsulation
    - Polymorphism
    - Inheritance
        But sometimes the last two are combined or confused.
- Abstraction -> decoupling behaviour from implementation details. UNIX file system is a great example. we use them but we don't know how they are implemented.
- Encapsulation -> hiding implementation details from misuse. It usually means controlling visibility of names (private variables)
- Polymorphism -> having many shapes. There are three main types:
    - Ad-hoc: typically found in function/operator overloading.
    - Parametric: k/a generic programming. (not in Go)
    - Subtype: subclass substituting for superclass.
- Prefer composition over inheritance.???
- Go offers abstraction & encapsulation; & polymorphism through interfaces. There is no inheritance in Go becoz there are no classes
- Go says -> inheritence is enough of a problem that are gonna do OOP without inheritance.

### Methods & Interfaces
- We can put methods on any user declared type.

### Composition
- it lets us build structs from other structs.
- composition is not inheritance.
- fields of the base struct get promoted into the derived struct. It means the fields of the base struct appear at the same level as the fields of the derived struct.
- If I have a base struct Pair & derived one PairWithLength, when PairWithLength goes to resolve a methods, it first checks if there a method defined with that name in PairWithLength, if yes then it uses that method, if not then it checks in Pair. So, derived one can either use the method of the base struct or override it.
- if a function takes a base struct as a parameter, it can not take a derived struct as a parameter, becoz they are not compatible.
- we can also use pointers to structs in composition.
