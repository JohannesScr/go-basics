package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("### Continue from Section 6")

	// let's look at the boolean type
	typeBool()
	// let's look at the int type
	typeInt()
	// let's look at the float type
	typeFloat()
	// let's look at the string type
	typeString()
	// let's look at the numeral system
	numeralSystems()
	// let's look at constants
	typeConstants()
	// let's look at iota
	typeIota()
	// let's look at bit shifting
	bitShifting()
	// exercises
	sec7ex1()
	sec7ex2()
	sec7ex3()
	sec7ex4()
	sec7ex5()
	sec7ex6()
}

// variable golbalX is of type bool
var globalX bool

func typeBool() {
	fmt.Println("\n### Bool Types")
	fmt.Println("denoting a system of algebraic notation to represent" +
		" logic propositions especially in computing and electronics\n" +
		"a binary variable, having two possible values called 'true' " +
		"or 'false'")
	fmt.Printf("globalX is of type %T and has a zero/default "+
		"value of %v\n", globalX, globalX)
	globalX = true
	fmt.Printf("Now we assigned globalX the value of 'true', %v\n",
		globalX)

	a := rand.Intn(5)
	b := rand.Intn(5)
	fmt.Printf("%v == %v: %v\n", a, b, a == b)
	fmt.Printf("%v <= %v: %v\n", a, b, a <= b)
	fmt.Printf("%v >= %v: %v\n", a, b, a >= b)
	fmt.Printf("%v != %v: %v\n", a, b, a != b)
}

func typeInt() {
	fmt.Println("\n### Int Types")
	fmt.Println("also referred to as whole numbers")
	fmt.Println("there are different options int, int8, int16, int32, " +
		"int64. and simply refers the size allocated to hte variable. in " +
		"in general you only need to use int, unless you know the size" +
		"the variable reange will be within.")
	fmt.Println("you also get uint and int. 'u' refers to unsigned or" +
		"positive, therefore unit = [0, 255] while int = [-128, 127]")
	fmt.Print("\n\tvar x int8\n\tx = 128 // will give an overflow " +
		"error\n\tx = -129 // will also give an overflow error\n")
	fmt.Println("'byte' is an alias for unit8")
	fmt.Println("'rune' or 'character' or 'utf-8' is an alias for unit32")
}

func typeFloat() {
	fmt.Println("\n### Float Types")
	fmt.Println("also referred to as real numbers")
	fmt.Println("there are different options float32, float64. also " +
		"refers to the size allocated to the variable")
}

func typeString() {
	fmt.Println("\n### String Types")
	fmt.Println("A string represents the set of string values. A string" +
		"is a (possibly empty) sequence of bytes. strings are immutable " +
		"(cannot be changed) once created. it is impossible to change the " +
		"contents of a string.")
	var doubleQuote string
	var backTick string
	doubleQuote = "Double Quote Hello World\n"
	backTick = `Back Tick Raw 
String 
Literal\n`
	fmt.Print(doubleQuote)
	fmt.Println(backTick)

	fmt.Println("You can however assign new value to the variable," +
		"although the string itself is immutable")
	doubleQuote = "New Assigned Double Quote Hello World\n"

	fmt.Println("\nLet's convert from string to slice of bytes")
	byteSlice := []byte(doubleQuote)
	fmt.Printf("%v\n is of type %T\n", byteSlice, byteSlice)

	//fmt.Println("print doubleQuote each character in UTF-8 format")
	//for i := 0; i < len(doubleQuote); i++ {
	//	fmt.Printf("%#U\n", doubleQuote[i])
	//}

	fmt.Println("index\tvalue\thex\tutf-8")
	for i, v := range doubleQuote {
		fmt.Printf("%d:\t%v\t%x\t%#U\n", i, v, v, v)
	}
	fmt.Println(doubleQuote)

	// strings, bytes, runes and characters in go
	// blog post
	// https://go.dev/blog/strings
}

func numeralSystems() {
	fmt.Println("### Numeral System")
	fmt.Println("as human we use the base 10 numeral system, consider")
	fmt.Println("\nBASE 10")
	fmt.Printf("10000\t1000\t100\t10\t1\n")
	fmt.Printf("10^4\t10^3\t10^2\t10^1\t10^0\n")
	fmt.Printf("0\t0\t1\t4\t2\t= 142\n")

	fmt.Println("\nBASE 2")
	fmt.Printf("32\t16\t8\t4\t2\t1\n")
	fmt.Printf("2^5\t2^4\t2^3\t2^2\t2^1\t2^0\n")
	fmt.Printf("1\t0\t1\t0\t1\t0\t= 42\n")

	fmt.Println("\nBASE 16")
	fmt.Println("Base 16 uses 10 characters {0-9} and another 6 " +
		"as {a, b, c, d, e, f}")
	fmt.Printf("32\t16\t8\t4\t2\t1\n")
	fmt.Printf("16^5\t16^4\t16^3\t16^2\t16^1\t16^0\n")
	fmt.Printf("0\t0\t0\t0\t0\t8\t= 8\n")
	fmt.Printf("0\t0\t0\t0\t0\ta\t= 10\n")
	fmt.Printf("0\t0\t0\t0\t0\tf\t= 15\n")
	fmt.Printf("0\t0\t0\t0\t1\t0\t= 16\n")
	fmt.Printf("0\t0\t0\t3\t8\tf\t= 911\n")

	fmt.Println("\nLet's do a quick conversion")
	s := "H"
	bs := []byte(s) // convert string to slice of type bytes
	n := bs[0]      // access the 0-th index
	fmt.Printf("s is %v: bs is %v: n is %v of type %T\n", s, bs, n, n)
	fmt.Printf("%b in bytes and %#X in hexadecimal\n", n, n)
}

func typeConstants() {
	fmt.Println("\n### Constants")

	const (
		a         = 42    // untyped constant
		b float32 = 42.78 // typed constant
		c         = "James Bond"
	)
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
}

const (
	_  = iota
	kb = 1 << (iota * 10) // 1 shifted (1 iota * 10 places)
	// => 10 places shifted over
	mb = 1 << (iota * 10) // 1 shifted (2 iota * 10 places)
	// => 20 places shifted over
	gb = 1 << (iota * 10) // 1 shifted (3 iota * 10 places)
	// => 30 places shifted over
	d = iota // 4
	e = iota // 5
)

func typeIota() {
	fmt.Println("\n### iota")
	fmt.Println("is a special pre-declared identifier")
	fmt.Println("if you need some counter")

	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", e, e)
}

func bitShifting() {
	fmt.Println("\n### Bit Shifting")

	x := 4
	fmt.Printf("%d\t%b\n", x, x)

	y := x << 1 // shifted 0100 -> 1000
	fmt.Printf("%d\t%b\n", y, y)

	z := x >> 1 // shifted 0100 -> 0010
	fmt.Printf("%d\t%b\n", z, z)

	kilobyte := 1024 // bytes
	megabyte := 1024 * kilobyte
	gigabyte := 1024 * megabyte

	fmt.Printf("kilobyte\t%d\t\t%b\n", kilobyte, kilobyte)
	fmt.Printf("megabyte\t%d\t\t%b\n", megabyte, megabyte)
	fmt.Printf("gigabyte\t%d\t%b\n", gigabyte, gigabyte)

	fmt.Printf("kilobyte\t%d\t\t%b\n", kb, kb)
	fmt.Printf("megabyte\t%d\t\t%b\n", mb, mb)
	fmt.Printf("gigabyte\t%d\t%b\n", gb, gb)
}

func sec7ex1() {
	fmt.Println("\n### Section 7 Exercise 1")
	/*
	Write a program that prints a number in decimal, binary, and hex
	*/
	x := 43
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

}
func sec7ex2() {
	fmt.Println("\n### Section 7 Exercise 2")
	/*
	Using the following operators, write expressions and assign their values to variables:
	==, <=, >=, !=, <, >
	Now print each of the variables.
	*/
	x := rand.Intn(5)
	y := rand.Intn(5)
	fmt.Printf("%v == %v is %v\n", x, y, x == y)
	fmt.Printf("%v <= %v is %v\n", x, y, x <= y)
	fmt.Printf("%v >= %v is %v\n", x, y, x >= y)
	fmt.Printf("%v != %v is %v\n", x, y, x != y)
	fmt.Printf("%v < %v is %v\n", x, y, x < y)
	fmt.Printf("%v > %v is %v\n", x, y, x > y)
}
func sec7ex3() {
	fmt.Println("\n### Section 7 Exercise 3")
	/*
	Create TYPED and UNTYPED constants. Print the values of the constants.
	*/
	const typed int = 5
	const untyped = 10
	fmt.Printf("typed constant %v of %T ", typed, typed)
	fmt.Printf("and untyped constant %v of type %T\n", untyped, untyped)
}
func sec7ex4() {
	fmt.Println("\n### Section 7 Exercise 4")
	/*
	Write a program that
		- assigns an int to a variable
		- prints that int in decimal, binary, and hex
		- shifts the bits of that int over 1 position to the left, and
		assigns that to a variable
		- prints that variable in decimal, binary, and hex
	*/
	x := 45
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	bitShift := byte(x) << 1
	fmt.Printf("%d\t%b\t%#x\n", bitShift, bitShift, bitShift)
}
func sec7ex5() {
	fmt.Println("\n### Section 7 Exercise 5")
	rawString := `raw string`
	fmt.Println(rawString)
}
func sec7ex6() {
	fmt.Println("\n### Section 7 Exercise 6")
	/*
	Using iota, create 4 constants for the NEXT 4 years.
	Print the constant values.
	*/
	const (
		thisYear = 2021
		year1 = thisYear + iota
		year2 = thisYear + iota
		year3 = thisYear + iota
		year4 = thisYear + iota
	)
	fmt.Println(year1, year2, year3, year4)
}


