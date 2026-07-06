package main

import "fmt"
import "math"
import "strconv"

func main() {
	fmt.Println("Hello, World!")

	// ----------------------------------- DATA TYPES -----------------------------------

	// Unsigned integers - only positive values
	var num uint = 1 // defaults to uint32 or uint64 based on the stored number
	var num2 uint8 = 1 // can store 0 to (2^8)-1 values
	var num3 uint32 = 1 // can store 0 to (2^32)-1 values
	var num4 uint64 = 1 // can store 0 to (2^64)-1 values

	// Signed integers - both negative and positive values (1st bit is sign bit)
	var num5 int = -1 // defaults to int32 or int64 based on the stored number
	var num6 int8 = -1 // can store -(2^7) to (2^7)-1 values
	var num7 int32 = 1 // can store -(2^31) to (2^31)-1 values
	var num8 int64 = 1 // can store -(2^63) to (2^63)-1 values

	// Decimal numbers
	var num9 float32 = 1.1
	var num10 float64 = -2.1

	// Byte - equivalent to int8
	var num11 byte = 8
	var char1 byte = 'a' // can be used to store a single character
	// var word byte = "abcd" // -> Error: cannot use "abcd" (untyped string constant) as byte value in variable declaration

	// Boolean - true or false
	var isBool bool = true

	// String
	var word1 string = "abcdefg" // single quotation mark does not work

	// nil - non-existent value, equivalent to null in javascript 
	

	fmt.Println(num, num2, num3, num4, num5, num6, num7, num8, num9, num10, num11, char1, isBool, word1)


	// ----------------------------------- VARIABLE DECLARATION AND ASSIGNMENT -----------------------------------

	// Explicit assignment - specifying the type explicitly
	// Use this when you want to declare a variable without any value (i.e., use its default value)
	var num12 int8 = 100
	var num13 int8 // uses default value (min value that can be stored), in this case 0

	fmt.Println(num12, num13)

	// Implicit assignment - type is determined by compiler based on stored value
	// Use this wherever possible
	num14 := 4 // type is infered as int or int32
	num15 := 2.3 // type is infered as float64

	fmt.Printf("%T", num14)
	fmt.Println()
	fmt.Printf("%T", num15)
	fmt.Println()

	// ----------------------------------- TYPE CASTING -----------------------------------

	num16:=uint(0)
	num17:=float32(0)

	fmt.Println(num16)
	fmt.Printf("%T",num16)
	fmt.Println()
	fmt.Println(num17)
	fmt.Printf("%T",num17)
	fmt.Println()

	// type casting can yield wrong results if trying to convert a value that is out of range for the final type
	num18 := -1000 // int type
	num19 := uint(num18) // uint cannot have negative values, its range starts from 0. This would not convert to positive 1000.

	fmt.Println(num19); // logs 18446744073709550616

	// ----------------------------------- FORMATTING STRINGS AND LOGGING OUTPUT -----------------------------------

	// Println - automatically adds a new line character at the end and spaces in between
	fmt.Println("Rohit","is",28,"years","old"); // Rohit is 28 years old 

	// Printf - can be used to create and print formatted strings with variables (not supported in Println)
	// Printf does not add a new line character at end
	name := "Rohit"
	age := 28
	fmt.Printf("%v is %v years old \n", name, age) // %v is used to interpolate value
	fmt.Printf("His age in binary is %b \n", age) // -> 11100 | %b is used to get the binary value of a number
	cgpa := 7.76543217765433
	fmt.Printf("Scientific notation for his CGPA is %e \n", cgpa) // -> 7.765432e+00 | %e is used to represent a floating value in its scientific notation
	fmt.Printf("CGPA in floating point is %f \n",cgpa) // -> 7.765432 | %f is used for float value rounded off to 6 decimal places
	fmt.Printf("CGPA rounded off to 2 decimal places is %.2f \n", cgpa) // ->7.77 | %.xf is used for float value rounded off to x decimal places
	university := "SRM IST"
	fmt.Printf("He did his graduation from %s \n", university) // %s is used to display string values
	percentage := 82
	fmt.Printf("He \"graduated\" with %v%% \n",percentage) // -> He "graduated" with 82% | use escape sequence (\) for any special character and %% for printing out "%"

	// Sprintf - can be used to create a formatted string that is not immediately printed but can be passed to a variable for future use
	graduationMessage := fmt.Sprintf("He \"graduated\" with %v%%",percentage)
	fmt.Println(graduationMessage)

	// ----------------------------------- ARITHMETIC OPERATORS -----------------------------------

	// operators present: + - * / ++ -- %
	// These operators need operands with same type to work

	num20 := uint(1000)
	num21 := 200

	//fmt.Println(num20+num21) // -> Error: invalid operation: num20 + num21 (mismatched types uint and int)
	fmt.Println(int(num20)+num21) // -> 1200

	// When using type conversions for arithmetic operators, we need to convert smaller type (less bits) to larger type (more bits), otherwise we risk going out-of-range for the smaller type
	num22 := uint8(10)
	num23 := 1000

	fmt.Println(num22 + uint8(num23)) // -> 242 | expected output was 1010, but since 1000 is out of uint8 range, we got wrong result
	fmt.Println(int(num22) + num23) // -> 1010

	// ----------------------------------- MATH PACKAGE -----------------------------------

	fmt.Println(math.Max(4,5)) // -> 5
	fmt.Println(math.Min(4,5)) // -> 4
	fmt.Println(math.Pow(4,2)) // -> 16
	fmt.Println(math.Sqrt(16)) // -> 4
	fmt.Println(math.Round(4.2345)) // -> 4
	fmt.Println(math.Ceil(4.2345)) // -> 5
	fmt.Println(math.Floor(4.6789)) // -> 
	
	// ----------------------------------- STRING CONVERT PACKAGE -----------------------------------

	// num := int("1234") // -> Error: cannot convert "1234" (untyped string constant) to type int

	num24, err1 := strconv.Atoi("1234") // converts to integer and returns two values: integer, error
	fmt.Println(num24, err1) // -> 1234 <nil>

	num25, err2 := strconv.Atoi("1234hello")
	fmt.Println(num25,err2) // -> 0 strconv.Atoi: parsing "1234hello": invalid syntax

	num26, err3 := strconv.ParseInt("1234", 10, 0) // same as Atoi (L140). 2nd parameter is the base we want to convert from and 3rd is the bit size of int which can be passed 0 as default (in which case it considers int64)
	fmt.Println(num26, err3) // -> 1234 <nil>

	num27, err4 := strconv.ParseInt("1234", 2, 0) // throws error as the 1234 is not a binary number for which we passed 2 as the base parameter
	fmt.Println(num27, err4) // -> 0 strconv.ParseInt: parsing "1234": invalid syntax

	num28, err5 := strconv.ParseInt("1110011", 2, 0) // works correctly as 1110011 is a binary number
	fmt.Println(num28, err5) // -> 115 <nil>

	// ----------------------------------- COMPARISON OPERATORS ----------------------------------- 

	// Operators present: < > <= >= == !=
	// Operand types need to be equal for these operators to work

	num29 := uint(8)
	num30 := 10

	// fmt.Println(num29 < num30) // -> Error: invalid operation: num29 < num30 (mismatched types uint and int)
	fmt.Println(num29 < uint(num30)) // -> true
	fmt.Println(num29 < 31) // -> true | this works without any type casting as when using a literal number, Go would try to convert the value automatically

	// ----------------------------------- CONDITION STATEMENTS ----------------------------------- 

	// If-Else

	num31 := 25

	if num31 > 30 {
		fmt.Println("num greater than 30")
	} else if num31 >= 20 {
		fmt.Println("num greater than 20")
	} else {
		fmt.Println("num less than 20")
	}

	// Classical Switch - break statements are not needed after a case statement

	switch num31 {
	case 35:
		fmt.Println("greater than 30")
	case 25: // only this will be caught
		fmt.Println("greater than 20")
	default:
		fmt.Println("less than 20")
	}

	// Naked Switch - do comparisons directly in case statements

	switch {
	case num31 > 30:
		fmt.Println("greater than 30")
	case num31 > 20: // only this will be caught
		fmt.Println("greater than 20")
	default:
		fmt.Println("less than 20")
	}

	// trigger fallthrough - go to the next case even if a previous case resolved to true

	switch num31 {
	case 35:
		fmt.Println("greater than 30")
		fallthrough
	case 25: // this will be caught
		fmt.Println("greater than 20")
		fallthrough
	case 20: // then this will be caught
		fmt.Println("less than 20")
	default:
		fmt.Println("undetermined")
	}

	// Multiple checks with the same case statement

	switch num31 {
	case 35, 34, 33, 32, 31, 30: // multiple values separated by commas
		fmt.Println("greater than 30")
	case 25, 24, 23, 22, 21, 20: // only this will be caught
		fmt.Println("greater than 20")
	default:
		fmt.Println("less than 20")
	}

	switch {
	case num31 > 35, num31 > 33, num31 > 30: // multiple checks separated by commas
		fmt.Println("greater than 30")
	case num31 > 25, num31 > 23, num31 > 20: // only this will be caught
		fmt.Println("greater than 20")
	default:
		fmt.Println("less than 20")
	}

	// ----------------------------------- LOOPS -----------------------------------

	// for loops - no parenthesis needed
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ",i)
	}
	fmt.Println()

	// for loops with range syntax
	for i := range 10 {
		fmt.Printf("%v ",i)
	}
	fmt.Println()

	// while loops - there is no while keyword in Go, we use "for" instead
	j := 0
	for j < 10 {
		fmt.Printf("%v ",j)
		j++;
	}
	fmt.Println()

	// ----------------------------------- STRING ACCESS & ITERATION -----------------------------------

	/* 
		Strings in Go are stored as a sequence of bytes
	 	In ASCII, 1 byte holds 1 character giving 256 unique character combinations
	 	In UTF-8, as much as 4 bytes can be used to hold 1 character (special characters, emojis, chinese letters, etc.)
		Thus, in UTF-8, when looping over a string, str[0] will return 1 byte data from the string which may or may not be equal to 1 character
		Hence, instead of using the normal for loop which will return 1 byte data with every iteration, range syntax should be used.
		The range syntax automatically takes care on handing different byte sized characters and returns 1 character each iteration.
	*/  
	str := "Go语😊"

	// index access on string - can fetch values at any index, but cannot update them
	fmt.Println(str[0]) // -> 71 | gives the integer (byte - uint8) representation of the character at index 0
	fmt.Println(string(str[0])) // -> G

	// str[0] = "d" // -> Error: cannot assign to str[0] (neither addressable nor a map index expression)

    // Normal for loop - iterates over 1 byte of data at a time
    for i := 0; i < len(str); i++ {
        fmt.Printf("%v ", string(str[i])) // -> G o è ¯  ð | does not produce correct result as last two characters use more than a single byte
    }

	fmt.Println()

    // Range loop - iterates over 1 rune or 1 character at a time
    for _, char := range str {
        fmt.Printf("%v ", string(char)) // -> G o 语 😊
    }

	fmt.Println()

	// _ can be used for values that you dont need. Range returns the index and character, however since index is not needed, an _ can be used to avoid compiler errors.
	for _, char := range str {
		fmt.Printf("%c ", char) // -> G o 语 😊 | %c is used to get the character value. works same as string(char)
    }
}