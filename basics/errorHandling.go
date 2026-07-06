package main

import "fmt"
import "errors"


// Using errors package

func concatenate(str1 string, str2 string) (string, error) {
	if str1 == "" && str2 == "" {
		return "", errors.New("Cannot concat empty strings")
	}

	return str1+" "+str2, nil
}

// Using the defer keyword and recover function

func divideByZeroHandler(){
	fmt.Println("division by zero")
	// recover function catches any errors thrown by a panic and returns it
	// it helps continue normal execution of a program without throwing errors in the console
	// it has to be called within a deffered function to catch runtime errors
	// returns nil if no errors are caught
	error := recover()
	if error != nil {
		fmt.Println(error) // -> runtime error: integer divide by zero
	} else {
		fmt.Println("no error")
	}
}

func divideNum(num1 int, num2 int) int {
	return num1/num2
}

func ErrorHandling(){
	/*
		Go is a compiled language so most errors are caught during compile time
		There are ways to handle run time errors like division by zero

		Such errors in Go are called panics
	*/

	// Handling errors with the errors package
	concatenated, err := concatenate("","")

	if(err != nil) {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println(concatenated)
	}

	// Handling errors with the defer keyword and recover function

	defer divideByZeroHandler() // defer keyword is used to run a piece of code at the end even after a panic occurs | this runs irrespective of a panic occurs or not
	result := divideNum(5,0) // -> panic: runtime error: integer divide by zero | this will not throw error in console if defered function calls recover()
	fmt.Println(result) // any piece of code after a panic is unreachable and does not run
}