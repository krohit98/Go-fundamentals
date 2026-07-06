package main

import "fmt"

func main(){

	// GENERICS

	/*
		In Go, generics are a language feature for parameterizing types.
		Generics let you write code once and use it with different types safely.
		Parameterized types can be used to create generic functions, generic structs, generic custom types, generic interfaces, etc.
	*/

	// using generic functions to perform the same operation over different types with the same function
	value1 := add(3, 4)
	value2 := add(3.2, 5.4)
	value3 := add(uint(3), uint(4))

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)

	value4 := subtract(10.9, 5.6)
	fmt.Println(value4)

	// Reading values from different map types
	map1 := map[string]int{"a":1, "b":2,"c":3}
	map2 := map[int]string{1:"a",2:"b",3:"c"}

	fmt.Println(getValues(map1))
	fmt.Println(getValues(map2))
	
	// using a generic type to create slices of different types and print their values using the same generic-type method 
	sl1 := GenericSlice[int]{1,2,3,4}
	sl2 := GenericSlice[string]{"a","b","c","d"}

	sl1.Print()
	sl2.Print()

	// using a generic struct to hold value of different types
	st1 := Box[int]{Value: 10}
	st2 := Box[string]{Value: "hello"}

	fmt.Println(st1.Value)
	fmt.Println(st2.Value)

	Pointers()
}

// GENERIC FUNCTIONS

// The add function can take up parameters of either int/float64/uint type
func add[T int|float64|uint](x T, y T) T {
	return x + y
}

// CONSTRAINT INTERFACES

// It is a better way of defining the constraint types a generic type should use
// It defines a reusable constraint for generic type parameters
// It is different than basic interfaces that declare method types (structsAndInterfaces.go)
type Number interface{
	int | float64 | uint
}



func subtract[T Number](x T, y T) T {
	return x - y
}

// creating a generic function to read values from a map of any type

/*
	- comparable is a built-in constraint interface that allows types that can be compared using == or !=. Map keys implement this.
	- "any" type allows the value to be of any type
*/
func getValues[K comparable, V any](mp map[K]V) []V {
	result := []V{}

	for _, val := range mp {
		result = append(result, val)
	}

	return result
}

// GENERIC TYPES

// creating a generic slice type that can hold a slice of any type
type GenericSlice[T any] []T // T is a type parameter representing a concrete type

// creating a generic type allows us to create methods for that type
func (g GenericSlice[T]) Print() {
	for _, val := range g {
		fmt.Println(val)
	}
}

// GENERIC STRUCTS

type Box[T any] struct {
	Value T

}
