package main

import "fmt"

// function accepts normal int parameter | creates a new variable and updates the new variable
func changeByValue(a int){
	a = 1000
	fmt.Println("inside change by value function:",a)
}

// function accepts a pointer parameter | uses a reference to the old variable and updates that
func changeByReference(a *int){
	*a = 1000
	fmt.Println("inside change by reference function:",*a)
}

// Pointers with structs

type Book struct {
	id int
	title string
}

// this setter function creates a new instance of book and only updates the new instance
func (b Book) setTitle(newTitle string) {
	b.title = newTitle
}

// this setter function uses a reference to the existing instance and updates the old instance
func (b *Book) setTitleWithPointer(newTitle string) {
	// dereference the instance and update the title
	// *b.title = newTitle // -> Error: invalid operation: cannot indirect b.title (variable of type string)
	(*b).title = newTitle // Parenthesis () is required as otherwise, compiler will try to destructure b.title instead of b
}

// Go is smart enough to automatically find and destructure the pointer variable
func (b *Book) setTtileWithPointerVariation(newTitle string) {
	b.title = newTitle // no need to destructure manually | only works with structs
}

// Advanced usecase - Pointer to a slice that holds pointers

func printValues(pointerSlice *[]*int) {
	// Slices have special formatting logic in the fmt package. When you print a slice (or a pointer to a slice), it automatically dereferences and formats the contents as [...], not just the address.
	fmt.Println(pointerSlice) // -> &[0x6aa07288e0b0 0x6aa07288e0b8 0x6aa07288e0c0 0x6aa07288e0c8] | The & at start denotes its a pointer

	values := *pointerSlice // destructure to get the actual slice
	fmt.Println(values) // -> [0x6aa07288e0b0 0x6aa07288e0b8 0x6aa07288e0c0 0x6aa07288e0c8]

	for _, value := range values {
		fmt.Println(*value) // destructure to print the actual int values
	}
}

func Pointers() {

	// Pointers are variables/identifiers used to store the hexadecimal memory address of another variable

	x := 100 // stores the value 100
	y := &x // stores the hexadecimal memory address of x
	z := &y // stores the hexadecimal memory address of y

	fmt.Println(x, y, z) // 100 0x399a6fe0e050 0x399a6fe10038

	// destructure (using *) to get the value at the memory address stored in a pointer

	fmt.Println(x, *y, **z) // -> 100 100 100
	
	// can update the actual value using destructure operator

	*y = 200

	fmt.Println(x, *y, **z) // -> 200 200 200

	**z = 300

	fmt.Println(x, *y, **z) // -> 300 300 300

	// Pass-by-value vs Pass-by-reference

	changeByValue(x) // -> inside change by value function: 1000
	fmt.Println("outside change by value function:",x) // -> outside change by value function: 300

	changeByReference(&x) // -> inside change by reference function: 1000
	fmt.Println("outside change by reference functon:",x) // -> outside change by reference functon: 1000

	fmt.Println(x, *y, **z) // -> 1000 1000 1000

	// Using pointers with structs

	book1 := Book{1, "old title"}
	fmt.Println(book1) // -> {1 old title}

	book1.setTitle("new title")
	fmt.Println(book1) // -> {1 old title} | this does not work as the setter creates a new instance of Book and updates its title 

	(&book1).setTitleWithPointer("new title") // pass a pointer to the instance. Parenthesis () is required as otherwise compiler will try to create a reference to book1.setTitleWithPointer(...)
	fmt.Println(book1) // -> {1 new title} | this works because the setter uses a reference to update the title of this instance

	// better way of using the setter function with pointer
	book1.setTtileWithPointerVariation("newest title") // no need to pass the pointer explicitly | Go can figure out to pass a pointer using the function definition
	fmt.Println(book1) // -> {1 newest title}

	// printing values from a pointr to a slice that holds pointers to int values
	a := 10
	b := 20
	c := 30
	d := 40

	pointerSlice := []*int{&a, &b, &c, &d}

	printValues(&pointerSlice) /* ->
		10
		20
		30
		40
	*/

	Goroutine1()
}