package main

import "fmt"

func main(){
	// ------------------------------------------------ ARRAYS ------------------------------------------------
	fmt.Println()
	fmt.Println("--------------- ARRAYS ---------------")
	fmt.Println()

	/*
		Fixed sized datastructure
		Can hold values of a single type
		The size of an array is defined at the time of initialization and it is considered a part of its type.
		The size of an array cannot be changed after initialization. 
		Cannot add or delete elements after initialization. Only mutate elements.
	*/
	
	var arr1 [2]int // [2]int is the type or arr1

	fmt.Printf("%T \n",arr1) // -> [2]int
	fmt.Println(arr1) // -> [0 0] | by default logs the default value of the datatype (0 for int)

	var arr2 [3]bool

	fmt.Println(arr2) // -> [false false false]

	// Array Literals

	arr3 := [3]int{1, 2, 3}

	fmt.Println(arr3) // -> [1 2 3]

	arr4 := [3][2]int{{1,2},{3,4},{5,6}} // nested array

	fmt.Printf("%T \n",arr4) // -> [3][2]int
	fmt.Println(arr4) 

	arr5 := [...][2]int{{1,2},{3,4}} // ... is used to automatically determine the size at compile time
	// arr6 := [...][...]int{{1,2},{3,4}} // -> Error: invalid use of [...] array (outside a composite literal) | Go only allows ... (length inference) for the outermost array dimension, not nested ones.

	fmt.Printf("%T \n",arr5) // -> [2][2]int
	fmt.Println(arr5)

	// arr7 := [3]int{1,2,"hello"} // -> Error: cannot use "hello" (untyped string constant) as int value in array or slice literal | cannot contain data of another type

	fmt.Println(arr5[0]) // -> [1 2] | index access is present

	// Mutating Array

	arr5[0][1] = 3
	fmt.Println(arr5[0]) // -> [1 3]

	// arr5[0] = {6,7} // -> Error: expected operand, found '{' | to replace a nested array a new array needs to be initialized
	arr5[0] = [2]int{6,7}

	fmt.Println(arr5[0]) // => [6 7]

	// Length of array

	fmt.Println(len(arr5), len(arr5[0])) // -> 2 2

	// Looping over array

	for i, val := range arr5 {
		fmt.Println(i, val)
	}

	for _, nested := range arr5 {
		for _, value := range nested {
			fmt.Println(value)
		}
	}

	// Passing arrays to functions

	arr6 := [1]string{"original"}
	mutateArray(arr6) // Arrays are passed by value by default

	fmt.Println(arr6) // -> [original] | Any mutation does not take place in the original array.


	// ------------------------------------------------ SLICE ------------------------------------------------
	fmt.Println()
	fmt.Println("--------------- SLICE ---------------")
	fmt.Println()

	/*
		It is a part of an array
		It contains three things: 
			Pointer - a pointer to the underlying array index where the slice starts
			Length - number of all the elements the slice currently holds
			Capacity - number of all the elements the slice can potentially hold. It is the length from pointer index to end of array
		More flexible than an array - elements can be appended

	*/

	arr7 := [5]int{1,2,3,4,5}

	sl1 := arr7[2:4] // slice from array index 2 to 4 (excluding 4) - [3 4]
	sl2 := arr7[:] // slice from start of array to end of array

	fmt.Println(sl1, sl2) // -> [3 4] [1 2 3 4 5]

	fmt.Println(sl1[0]) // -> 3 | index access on a slice provides values according to slice index and not array index

	sl2[0] = 100

	fmt.Println(sl2, arr7) // -> [100 2 3 4 5] [100 2 3 4 5] | Modifying a slice also modifies the underlying array and vice versa

	// length and capacity

	fmt.Println(len(sl1), cap(sl1)) // -> 2 3

	// slice of a slice - can expand till capacity

	sl3 := sl1[:3] // creates a new slice starting from the start of the previous slice and expanding till capacity of the previous slice
	
	fmt.Println(sl3) // -> [3 4 5]

	// Creating slice without an array

	sl4 := []string{"hello","world"}

	fmt.Println(sl4)
	fmt.Printf("%T \n", sl4) // -> []string | this is the datatype for a slice

	// Appending to slice

	fmt.Println(sl4, len(sl4), cap(sl4))

	for x := 0; x <= 10; x++ {
		sl4 = append(sl4, "rohit")
		fmt.Println(sl4, len(sl4), cap(sl4))
	}
	/*
		output for the above loop is:

		[hello world] 2 2
		[hello world rohit] 3 4
		[hello world rohit rohit] 4 4
		[hello world rohit rohit rohit] 5 8
		[hello world rohit rohit rohit rohit] 6 8
		[hello world rohit rohit rohit rohit rohit] 7 8
		[hello world rohit rohit rohit rohit rohit rohit] 8 8
		[hello world rohit rohit rohit rohit rohit rohit rohit] 9 16
		[hello world rohit rohit rohit rohit rohit rohit rohit rohit] 10 16
		[hello world rohit rohit rohit rohit rohit rohit rohit rohit rohit] 11 16
		[hello world rohit rohit rohit rohit rohit rohit rohit rohit rohit rohit] 12 16
		[hello world rohit rohit rohit rohit rohit rohit rohit rohit rohit rohit rohit] 13 16

		Initially, the slice was created with a length and capacity of 2 each as it was initialized with 2 elements.
		When appending to a slice, if the slice is not full upto its capacity (length < capacity), the new element will be directly appended and only length will increase.
		However, when the slice is filled upto its max capacity (length = capacity), a new slice is created with double capacity and the new element is appended.
		The capacity is doubled to avoid increasing the capacity on each append.
		Note: append() always returns a new slice
	*/

	// Creating slice using make function - used to dynamically create slices with specified size (and capacity) without creating an array
	// make(type, length, capacity?) - capacity is optional. If not passed, the capacity is set to the length
	
	sl5 := make([]int, 10, 20)
	fmt.Println(sl5) // -> [0 0 0 0 0 0 0 0 0 0] | creates a slice of length 10 and capacity 20 with default int values (0)

	// Looping over slice

	sl6 := []string{"hello","world","marvel"}

	for i,value := range sl6 {
		fmt.Println(i, value)
	}

	// Passing slice to a function

	sl7 := []string{"original"}
	mutateSlice(sl7) // slices are passed by reference by default

	fmt.Println(sl7) // -> [updated] | any mutation updates the original slice


	// ------------------------------------------------ MAP ------------------------------------------------
	fmt.Println()
	fmt.Println("--------------- MAPS ---------------")
	fmt.Println()

	// Explicit assignment - with var syntax

	var map1 map[string]int = map[string]int{"a":1} // map[type-of-property]type-of-value
	fmt.Println(map1) // -> map[a:1]

	// Implicit assignment

	map2 := map[string]int{"b":2}
	fmt.Println(map2) // -> map[b:2]

	// map3 := map[string]int{"c":"c"} // Error: cannot use "c" (untyped string constant) as int value in map literal
	
	map4 := map[string]int{} // need to put empty curly braces to initialize empty
	fmt.Println(map4) // -> map[]

	// creating map using make function

	map5 := make(map[string]int) // do not need empty {} here
	fmt.Println(map5) // -> map[]

	// using complex types for value
	
	map6 := map[string][]int{"c":{1,2,3}} // type of value is slice here (slice datatype is []int here - slice of int values)
	fmt.Println(map6) // -> map[c:[1 2 3]]

	// Adding a new key

	map6["a"] = []int{10, 20, 30}
	fmt.Println(map6) // -> map[a:[10 20 30] c:[1 2 3]]

	map5["d"]++ // a key that does not exist in a map can be thought of as existing with the default value of its type-of-value
	fmt.Println(map5) // -> map[d:1] | this works because initially map5[d:0] (any non-existent key will have the same)

	// Updating value for existing key

	map6["c"] = []int{15,25,35}
	fmt.Println(map6) // -> map[a:[10 20 30] c:[15 25 35]]

	// Deleting a key

	delete(map6, "c") // delete does not return a new map unlike append in slice
	fmt.Println(map6) // -> map[a:[10 20 30]]

	// Accessing values

	value, isPresent := map6["a"] // returns the value of key and a boolean specifying if the key is present in the map.
	fmt.Println(value, isPresent) // -> [10 20 30] true

	value2, isPresent2 := map6["c"]
	fmt.Println(value2, isPresent2) // -> [] false

	Functions()
}

func mutateArray(arr [1]string){
	arr[0] = "updated" // does not mutate the original array
	fmt.Println(arr) // -> [updated]
}
 
func mutateSlice(sl []string){
	sl[0] = "updated" // mutates the original slice
	fmt.Println(sl) // -> [updated]
}