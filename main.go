package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(message string) {
	fmt.Println(message)
}

func add() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

type Meth struct {
	value int
}

func (v Meth) square() int {
	return v.value * v.value
}

func main() {

	// For loop
	// Go has only one looping construct, the for loop.
	for i := 15; i > 0; i-- {
		fmt.Println(i)
	}

	// 	For is Go's "while"
	// At that point you can drop the semicolons: C's while is spelled for in Go.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 	Forever
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	// for {
	// 	fmt.Println("he")
	// }

	// 	If else
	// Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( )
	// but the braces { } are required.
	const ifVar int = 3
	if ifVar < 5 {
		fmt.Println("ifVar is smaller than 5")
	} else {
		fmt.Println("ifVar is greater than 5")
	}

	// Switch
	// A switch statement is a shorter way to write a sequence of if - else statements.
	// It runs the first case whose value is equal to the condition expression.

	// Note - Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println(os)

	default:
		fmt.Println("default")
	}

	// 	Switch with no condition
	// Switch without a condition is the same as switch true.

	// This construct can be a clean way to write long if-then-else chains.
	switchT := time.Now()

	switch {
	case switchT.Hour() < 12:
		fmt.Println("good morning")

	case switchT.Hour() < 17:
		fmt.Println("Good afternoon")

	default:
		fmt.Println("good evening")
	}

	// 	Defer
	// A defer statement defers the execution of a function until the surrounding function returns.

	// The deferred call's arguments are evaluated immediately, but the function call is not
	// executed until the surrounding function returns.

	defer fmt.Println("morning")

	fmt.Println("Good")

	// 	Stacking defers
	// Deferred function calls are pushed onto a stack. When a function returns,
	// its deferred calls are executed in last-in-first-out order.
	fmt.Println("Deffer counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Deffer counting done!")

	// 	Pointers in Go (Golang):
	// Definition: Pointers are variables that store memory addresses of other variables.
	// Declaration: Declare a pointer using * followed by the type it points to, e.g., var ptr *int.
	// Initialization: Initialize a pointer with the address of a variable using &, e.g., var ptr *int = &num.
	// Dereferencing: Access the value at the memory address pointed to by a pointer using *ptr.
	// Zero Value: An uninitialized pointer has a zero value (nil), indicating it points to nothing.
	// Changing Values: Modify a variable's value through a pointer, e.g., *ptr = 100.
	// Passing to Functions: Pass variables by reference to functions for in-place modification.

	i, j := 42, 2701

	p := &i

	q := &j

	*p = 56

	fmt.Println(*p, *q)
	fmt.Println(p)

	// --------------------------------------------------------------------------------------------------------------
	// 	Structs
	// A struct is a collection of fields.
	// Struct fields are accessed using a dot.

	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference.
	type Node struct {
		value int
		next  int
	}

	nodeN := Node{2, 5}
	nodeNPtr := &nodeN
	fmt.Println(nodeN)
	fmt.Println(nodeN.value)
	fmt.Println(nodeN.next)
	fmt.Println(nodeNPtr)

	// --------------------------------------------------------------------------------------------------------------

	// Arrays
	// The type [n]T is an array of n values of type T.
	// Note - An array's length is part of its type, so arrays cannot be resized.

	var arr [10]int
	arr[5] = 1
	fmt.Println(arr)

	// --------------------------------------------------------------------------------------------------------------

	// 	Slices
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the
	// elements of an array. In practice, slices are much more common than arrays.

	// The type []T is a slice with elements of type T.

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon: sliceA[low:high]
	// This selects a half-open range which includes the first element, but excludes the last one.

	// A slice has both a length and a capacity.

	// The zero value of a slice is nil.

	slicesVar := []int{2, 3, 4}

	slicesVar[0] = 66
	fmt.Println(slicesVar[0:2])
	fmt.Println(len(slicesVar))
	fmt.Println(cap(slicesVar))

	// slice with built-in make function

	sliceMake := make([]int, 0, 5)

	fmt.Println(sliceMake)

	// --------------------------------------------------------------------------------------------------------------

	// 	Maps
	// A map maps keys to values.

	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

	// The make function returns a map of the given type, initialized and ready for use.

	mapP := make(map[string]Node)
	mapP["default"] = Node{5, 8}

	fmt.Println(mapP["default"])

	elem, ok := mapP["default"]
	fmt.Println(elem, ok)

	delete(mapP, "default")

	elem1, ok1 := mapP["default"]
	fmt.Println(elem1, ok1)

	// --------------------------------------------------------------------------------------------------------------

	// 	Function values
	// Functions are values too. They can be passed around just like other values.

	// Function values may be used as function arguments and return values.

	print("hello world")

	printValue := func(message string) {
		fmt.Println(message)
	}

	printValue("hello world")

	// --------------------------------------------------------------------------------------------------------------

	// Function closures
	// Go functions may be closures. A closure is a function value that references variables from
	// outside its body. The function may access and assign to the referenced variables; in this sense the
	// function is "bound" to the variables.

	pos, neg := add(), add()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// --------------------------------------------------------------------------------------------------------------

	// 	Methods
	// Go does not have classes. However, you can define methods on types.

	// A method is a function with a special receiver argument.

	// The receiver appears in its own argument list between the func keyword and the method name.

	// Remember: a method is just a function with a receiver argument.
	methValue := Meth{5}
	fmt.Println(methValue.square())

	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------------------------------

}
