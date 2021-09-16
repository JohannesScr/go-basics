package main

import "fmt"

func main() {
	fmt.Println("In the beninging...")
	n, err := fmt.Println(12.4, true, false)
	m, _ := fmt.Println("Hi there", 42) // _ means return will not be used
	fmt.Println("print ", n, " and ", m, " bytes with ", err, " errors")

	// declaring variables
	declareVariable()
	// declaring variable's defaults
	declareVariableDefaults()
	// declare my own type
	declareOwnType()
	// type conversion
	typeConversion()

	// Section 5 Exercise 1
	sec5ex1()

	// Section 5 Exercise 2
	sec5ex2()

	// Section 5 Exercise 3
	sec5ex3()

	// Section 5 Exercise 4
	sec5ex4()

	// Section 5 Exercise 5
	sec5ex5()

	// Section 5 Exercise 6
	sec5ex6()
}

func declareVariable() {
	fmt.Println("\n### Go Variables Declaration")
	// short declaration
	x := 42
	// explicit declaration
	var myFloat float64
	// assign a value to float64 variable
	myFloat = 12.4
	// explicit declaration with assignment
	var myString string = "this is my string"
	// although the variable type is not needed since the string already
	//implies the type, it can be added
	fmt.Println(x, myFloat, myString)
}

func declareVariableDefaults() {
	var a int
	var b float32
	var c string
	var d bool

	fmt.Println("\n### Go Variable Declaration Defaults")
	fmt.Printf("%T default is %v\n", a, a)
	fmt.Printf("%T default is %v\n", b, b)
	fmt.Printf("%T default is %v\n", c, c)
	fmt.Printf("%T default is %v\n", d, d)
}

// normal golang
var a int

// declare my own type
type hotdog int  //  declare a type hotdog which is an int
var b hotdog

func declareOwnType() {
	fmt.Println("\n### Declare Your Own Types")
	a = 42
	b = 42
	fmt.Printf("%v is a %T\n", a, a)
	fmt.Printf("%v is a %T\n", b, b)
	fmt.Println("Therefore we cannot assign a = b since their types " +
		"are different")
}

func typeConversion() {
	fmt.Println("\n### Type Conversion")
	var a int
	var b float32 = 12.3456
	fmt.Printf("a: %v is a %T\n", a, a)
	fmt.Printf("b: %v is a %T\n", b, b)
	fmt.Println("Conversion in go is the same as casting in other " +
		"programming languages. So we convert the float to a int and" +
		"vice versa")
	a = int(b)
	b = float32(a)
	fmt.Printf("a = int(b): %v is a %T\n", a, a)
	fmt.Printf("b = float32(a): %v is a %T\n", b, b)
	fmt.Println("So we need to convert the type of the variable being " +
		"assigned to the type of the variable being assigned to. Be " +
		"careful of data loss with type conversion")
}

func sec5ex1() {
	fmt.Println("\n### Section 5 Exercise 1")
	/*
	Using the short declaration operator, ASSIGN these VALUES to
	VARIABLES with the IDENTIFIERS “x” and “y” and “z”
		- 42
		- “James Bond”
		- true
	Now print the values stored in those variables using
		- a single print statement
		- multiple print statements
	*/
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}

var x int
var y string
var z bool
func sec5ex2() {
	fmt.Println("\n### Section 5 Exercise 2")
	/*
	Use var to DECLARE three VARIABLES. The variables should have package
	level scope. Do not assign VALUES to the variables. Use the following
	IDENTIFIERS for the variables and make sure the variables are of the
	following TYPE (meaning they can store VALUES of that TYPE).
		- identifier “x” type int
		- identifier “y” type string
		- identifier “z” type bool
	in func main
		- print out the values for each identifier
		- The compiler assigned values to the variables. What are these
		values called?
	*/
	fmt.Printf("%T is %v\n", x, x)
	fmt.Printf("%T is %v\n", y, y)
	fmt.Printf("%T is %v\n", z, z)
	fmt.Println("These default values assigned to a variable when it " +
		"is assigned is called the ZERO values")
}

func sec5ex3() {
	fmt.Println("\n### Section 5 Exercise 3")
	/*
	Using the code from the previous exercise,
	At the package level scope, assign the following values to the three
	variables
		- for x assign 42
		- for y assign “James Bond”
		- for z assign true
	in func main
		- use fmt.Sprintf to print all the VALUES to one single string.
		ASSIGN the returned value of TYPE string using the short declaration
		operator to a VARIABLE with the IDENTIFIER “s”
		- print out the value stored by variable “s”
	*/
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}

func sec5ex4() {
	fmt.Println("\n### Section 5 Exercise 4")
	/*
	For this exercise
		- Create your own type. Have the underlying type be an int.
		- Create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
		in func main
			- print out the value of the variable “x”
			- print out the type of the variable “x”
			- assign 42 to the VARIABLE “x” using the “=” OPERATOR
			- print out the value of the variable “x”
	*/
	type chairman int
	var x chairman

	fmt.Printf("%v is %T\n", x, x)
	x = 42
	fmt.Printf("%v is %T\n", x, x)
}

func sec5ex5() {
	fmt.Println("\n### Section 5 Exercise 5")
	/*
	Building on the code from the previous example
		- at the package level scope, using the “var” keyword, create a
		VARIABLE with the IDENTIFIER “y”. The variable should be of the
		UNDERLYING TYPE of your custom TYPE “x”
		- in func main
			- this should already be done
				- print out the value of the variable “x”
				- print out the type of the variable “x”
				- assign your own VALUE to the VARIABLE “x” using
				the “=” OPERATOR
				- print out the value of the variable “x”
			- now do this
				- now use CONVERSION to convert the TYPE of the VALUE
				stored in “x” to the UNDERLYING TYPE
					- then use the “=” operator to ASSIGN that value to “y”
					- print out the value stored in “y”
					- print out the type of “y”
	*/
	type chairman int
	var x chairman
	x = 42
	var y int

	fmt.Printf("%T is %v\n", x, x)
	y = int(x)
	fmt.Printf("%T is %v\n", y, y)
}

func sec5ex6() {
	fmt.Println("\n### Section 5 Exercise 6")
	/*

	*/

}
