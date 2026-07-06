package main

import "fmt"
import "math"

/*
	STRUCTS

	Custom types that can have fields of various datatypes and methods. These types can be used to create instances
	much like classes in other languages
*/

type Person struct {
	name string
	age int
}

// embedded struct - struct inside another struct
type Organisation struct {
	name string
	employees []Person // slice of a struct in this case
}

type WelcomeMessage struct {
	// function type written inside a struct is instance specific, i.e., implementation can be different per instance
	getWelcomeMessage func(string) string
}

// method of a struct - static definition, does not change per instance
func (p Person) greet() string {
	return "Hello, "+p.name+". Greetings!"
}

// Below syntax does not add a function as a struct method
func getBirthYear(p Person) int {
	return 2026 - p.age
}

// setter function to set a field value inside a struct - a copy of the struct instance is passed here
func (p Person) setName(newName string) {
	p.name = newName
	fmt.Println("inside setter", p)
}

// function that accepts a struct as a parameter
func getName(p Person) string{
	return p.name
}

// function that returns a struct as a parameter
func getPerson(name string, age int) Person{
	p := Person{name, age}
	return p
}

/* 
	INTERFACES

	An interface is like a wrapper over a struct that abstracts away information about the underlying struct and 
	provides a set of functions that are implemented as methods of the underlying structs

	All the functions declared in an interface must be implemented by a struct in order for the interface to wrap
	over the struct

	Instances of interfaces are not created.

	Helps implement Polymorphism as the functions can have different implementations across different structs
*/

type Shape interface{
	getArea() float64
	getPerimeter() float64
}

// first struct that implements the functions from the interface as its methods
type Triangle struct{
	a float64
	b float64
	c float64
}

func (t Triangle) getArea() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s*(s-t.a)*(s-t.b)*(s-t.c))
}

func (t Triangle) getPerimeter() float64 {
	return t.a + t.b + t.c
}

// second struct that implements the functions from the interface as its methods
type Rectangle struct{
	length float64
	breadth float64
}

func (r Rectangle) getArea() float64 {
	return r.length * r.breadth
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// function that takes in an interface type as parameter
// this makes the function much more flexible, as any struct type that implements the interface can be used here
func isEvenParameter(shape Shape) bool {
	return int(shape.getPerimeter()) % 2 == 0
}

func StructsAndInterfaces(){
	// ------------------------------------------------ STRUCTS ------------------------------------------------

	// with explicit assignment
	var p0 Person = Person{"Rohit", 28}
	fmt.Println(p0) // -> {Rohit 28}

	// with implicit assignment
	p1 := Person{} // Empty structs are initialised with the default values of the data types of its fields
	fmt.Println(p1) // -> { 0} | Prints p1 with default values - empty string for string and 0 for int

	p2 := Person{"Rohit", 28}
	fmt.Println(p2) // -> {Rohit 28}

	// cannot be initialised in the wrong order (datatype wise)
	// p3 := Person{28, "Rohit"} // -> Error: cannot use 28 (untyped int constant) as string value in struct literal

	// can be initialised in any order if the field names are specified
	p3 := Person{age:28, name:"Rohit"}
	fmt.Println(p3) // -> {Rohit 28}

	// can be initialised with values for one or more fields
	p4 := Person{age:30}
	fmt.Println(p4) // -> { 30}

	// has access to individual fields
	p4.name = "Mohit" // can be modified
	fmt.Println(p4) // -> {Mohit 30}
	fmt.Println(p4.age) // -> 30 | can be read

	// Passing struct to function
	name := getName(p4)
	fmt.Println(name) // Mohit

	// Getting struct from function
	p5 := getPerson("Rohit", 28)
	fmt.Println(p5) // -> {Rohit 28}

	// defining instance specific function of a struct

	welcomeMessage1 := WelcomeMessage{}
	welcomeMessage1.getWelcomeMessage = func(name string) string {
		return "Hello, "+name
	}
	fmt.Println(welcomeMessage1.getWelcomeMessage(p5.name)) // -> Hello, Rohit

	welcomeMessage2 := WelcomeMessage{}
	welcomeMessage2.getWelcomeMessage = func(txt string) string {
		return "Hello there, "+txt
	}
	fmt.Println(welcomeMessage2.getWelcomeMessage("how are you?")) // -> Hello there, how are you?

	// calling struct method

	fmt.Println(p5.greet()) // -> Hello, Rohit. Greetings!
	// fmt.Println(p5.getBirthYear()) // -> Error: p5.getBirthYear undefined (type Person has no field or method getBirthYear)

	// calling setter method
	p5.setName("Sheldon") // -> inside setter {Sheldon 28}
	fmt.Println(p5) // -> {Rohit 28} | a copy of the original struct instance is passed, thus original instance is not modified

	// Creating and using embedded structs

	// employeeSlice1 := []Person{p4, p5, Person{"Raj", 34}} // -> Warning: redundant type from array, slice, or map composite literal
	employeeSlice1 := []Person{p4, p5, {"Raj", 34}}
	org1 := Organisation{name: "Google", employees:employeeSlice1}

	fmt.Println(org1) // -> {Google [{Mohit 30} {Rohit 28} {Raj 34}]}
	fmt.Println(org1.employees[1].name) // -> Rohit

	org1.employees[0].age = 25
	fmt.Println(org1.employees[0]) // -> {Mohit 25} | embedded instance can be modified
	fmt.Println(p4) // -> {Mohit 30} | original struct instance is not modified

	// ------------------------------------------------ INTERFACES ------------------------------------------------

	// shape1 := Shape{a:15, b:15, c:20} // This is wrong as an instance of an interface is not created
	// shape1 := Triangle{a:15, b:15, c:20} // This is defining a regular instance of a struct, nothing related to interface

	// This creates an instance of Triangle (struct) which is stored in a variable of type Shape (static type - before runtime).
	var shape1 Shape = Triangle{a:15, b:15, c:20} // Hides away the data members of the underlying struct. Only methods declared in the interface are accessible.

	fmt.Printf("%T \n",shape1) // -> main.Triangle | At runtime, the dynamic type of shape1 is Triangle

	// fmt.Println(shape1.a) // -> Error: shape1.a undefined (type Shape has no field or method a)
	fmt.Println("Triangle Area",shape1.getArea())
	fmt.Println("Triangle Perimeter", shape1.getPerimeter())

	shape1 = Rectangle{length:5, breadth:10}

	fmt.Printf("%T \n",shape1) // -> main.Rectangle | At runtime, the dynamic type of shape1 is Rectangle

	// fmt.Println(shape1.length) // -> Error: shape1.length undefined (type Shape has no field or method length)
	fmt.Println("Rectangle Area",shape1.getArea())
	fmt.Println("Rectangle Perimeter", shape1.getPerimeter())

	// Slice of interface types
	var shapes []Shape = []Shape{Triangle{a:15, b:15, c:20}, Rectangle{length:5, breadth:10}}

	totalPerimeters := 0

	for _, shape := range shapes {
		totalPerimeters += int(shape.getPerimeter())
	}

	fmt.Println("Total Perimeters", totalPerimeters)

	// functions that take in interface type as parameter can be passed any struct that implements that interface
	fmt.Println(isEvenParameter(shape1)) // -> True
	fmt.Println(isEvenParameter(Triangle{a:15, b:15, c:20})) // -> True
	fmt.Println(isEvenParameter(Rectangle{length:5, breadth:10})) // -> True

	ErrorHandling()
}