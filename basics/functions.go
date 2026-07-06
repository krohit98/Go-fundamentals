package main

import "fmt"

func Functions(){
	fmt.Println()
	fmt.Println("--------------- FUNCTIONS ---------------")
	fmt.Println()

	value1 := add(1,2)
	fmt.Println(value1) // -> 3

	// value2 := mul(3,4) // -> Error: assignment mismatch: 1 variable but mul returns 2 values
	// if multiple values are returned, there must be equal variables to catch them
	value2, text := mul(3,4)
	fmt.Println(value2, text) // -> 12 multiplied

	// Passing function into another function as parameter

	// anonymous function
	double := func (num int) int {
		return 2 * num
	}

	// passing separately defined anonymous function
	value3 := callFunc(10, double) // we can also pass a normal function which is defined separately
	fmt.Println(value3) // -> 20

	// passing anonymous function defined directly in function call
	value4 := callFunc(10, func (num int) int {
		return 3 * num
	})
	fmt.Println(value4) // -> 30

	// getting function from another function 

	quadruple := getFunc(4)
	value5 := quadruple(10)
	fmt.Println(value5) // -> 40

	fiveTimes := getFunc(5)
	value6 := fiveTimes(10)
	fmt.Println(value6) // -> 50

	// passing variable number of parameters to a function (variadic function)

	sum1 := sum(10,10,10)
	fmt.Println(sum1) // -> 30

	// parameters can also be passed as a slice but it needs to be destructured into individual elements using ... syntax.
	// sum2 := sum([]int{20,20,20}) // -> Error: cannot use []int{…} (value of type []int) as int value in argument to sum
	sum2 := sum([]int{20,20,20}...) // this only works with variadic functions
	fmt.Println(sum2) // -> 60

	// calling functions with named return values
	difference := subtract(10,5)
	fmt.Println(difference) // -> 5

	quotient, operatorText := divide(14,7)
	fmt.Println(quotient, operatorText) // -> 2 division

	StructsAndInterfaces()
}

// Simple function - func function-name ( param-name param-type ) return-type { function-body }
func add(num1 int, num2 int) int {
	return num1+num2;
}

// return multiple values - put the return types in paranthesis
func mul(num1 int, num2 int)(int, string){
	return num1*num2, "multiplied"
}

// accepting functions as parameters
func callFunc(num int, callback func(int)int) int {
	value := callback(num)
	return value
}

// returning function from another function
func getFunc(num int) func(int) int {
	multiply := func (num2 int) int {
		return num2 * num
	}
	return multiply
}

// passing variable number of parameters (variadic function)
func sum(nums ...int) int { // type of this argument is a slice (nums []int) | internally it converts the passed params into a slice
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

// named return values
// the variable to return is defined along with type
// return keyword will return this variable by default
func subtract(num1 int, num2 int)(diff int){
	diff = num1 - num2
	return
}

func divide(num1 int, num2 int)(quotient int, operation string){
	quotient = num1/num2
	operation = "division"

	return
}
