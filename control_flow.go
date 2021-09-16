package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("### Control Flow ###")
	fmt.Println("control flow it is how the computer reads the program")
	fmt.Println("\tsequential")
	fmt.Println("\tloop / iterative")
	fmt.Println("\tconditional")

	// let's look at loops
	loops()
	// let's look at if and else
	ifElse()
	// loop condition modulus
	loopConditionModulus()
	// let's look at the switch statement
	switchStatement()
	// let's look at conditional and/or
	andOr()
	// hands-n exercises section 9
	section9ex1()
	section9ex2()
	section9ex3()
	section9ex4()
	section9ex5()
	section9ex6()
	section9ex7()
	section9ex8()
	section9ex9()
}

func loops() {
	fmt.Println("\n### Loops")
	fmt.Println("Basic for loop structure")
	fmt.Println("for init; condition; post {\n\t// code here\n}")
	// option one
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	// option two
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		for j := 2; j > 0; j-- {
			fmt.Println(i, j)
		}
	}
	// option three
	fmt.Println("Single Condition")
	x := 1
	for x < 10 {
		fmt.Println(x)
		x *= 2
	}
	// option four
	fmt.Println("With For Clause, the 'normal' for loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// option five
	fmt.Println("Range Clause")
	y := "hello"
	for index, value := range y {
		fmt.Println(index, value)
	}
	// option six
	fmt.Println("The 'break' statement")
	x = 1
	for {
		if x > 4 {
			break
		}
		fmt.Println(x)
		x++
	}
	// option seven
	fmt.Println("The 'continue' statement")
	x = 1
	for {
		if x < 5 {
			x++
			continue
		}
		fmt.Println("only print 5", x)
		break
	}
	// mini challenge
	fmt.Printf("int\tunicode\tcharacter\thex\n")
	for i := 33; i < 123; i++ {
		fmt.Printf("%d\t%U\t%q\t\t%#x\n", i, i, i, i)
	}
}

func ifElse() {
	fmt.Print("\n### If {} ")
	fmt.Println("Else {}")
	if true {
		fmt.Println("true")
	}
	if false {
		fmt.Println("false")
	}
	if !true {
		fmt.Println("!true")
	}
	if !false {
		fmt.Println("!false")
	}
	if 2 == 2 {
		fmt.Println("2 == 2")
	}
	if !(2 != 2) {
		fmt.Println("!(2 != 2)")
	}
	// initialise if statement
	y := rand.Intn(5)
	if x := 42; x != y {
		fmt.Printf("initialise x = %d; check %d != %d\n", x, x, y)
	}
	// if else
	x := rand.Intn(5)
	if x == y {
		fmt.Printf("%d == %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d > %d\n", x, y)
	} else {
		fmt.Printf("%d < %d\n", x, y)
	}
}

func loopConditionModulus() {
	for i := 0; i < 30; i++ {
		if (i % 3) == 0 {
			fmt.Printf("%d is divisable by 3\n", i)
		}
	}
}

func switchEquals(switchCase int) {
	fmt.Println("Strict Equal with Default Case")
	switch switchCase {
	case 2, 4:
		fmt.Println("2 or 4")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Default Case")
	}
}

func switchCondition(switchCase int) {
	fmt.Println("Strict With Condition Default Case")
	switch { // defaults to true
	case switchCase == 2: // then it checks the next conditions
		fmt.Println("2")
	case switchCase > 2:
		fmt.Println("3")
	default:
		fmt.Println("Default Case")
	}
}

func switchFallThrough(switchCase int) {
	fmt.Println("Strict Equal with Fall Through Default Case")
	switch switchCase {
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
		fallthrough
	default:
		fmt.Println("Default Case")
	}
}

func switchSimpleStatement(switchCase int) {
	fmt.Println("\n### Switch Simple Statement")
	switch 3 == switchCase {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}

func switchStatement() {
	fmt.Println("\n### Switch Statement")
	switchEquals(1)
	switchEquals(2)
	switchEquals(3)
	switchEquals(4)

	switchCondition(1)
	switchCondition(2)
	switchCondition(3)

	switchFallThrough(1)
	switchFallThrough(2)
	switchFallThrough(3)

	switchSimpleStatement(4)
}

func andOr() {
	fmt.Println("\n### Conditional And and Or")
	fmt.Printf("and\t%v && %v\t= %v\n", true, true, true == true)
	fmt.Printf("and\t%v && %v\t= %v\n", true, false, true == false)
	fmt.Printf("or\t%v || %v\t= %v\n", true, true, true || true)
	fmt.Printf("or\t%v || %v\t= %v\n", true, false, true || false)
	fmt.Printf("not\t!%v\t\t= %v\n", true, !true)
	fmt.Printf("not\t!%v\t\t= %v\n", false, !false)
}

func section9ex1() {
	fmt.Println("\nSection 9 Exercise 1")
	/*
		Print every number from 1 to 10,000
	*/
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
func section9ex2() {
	fmt.Println("\nSection 9 Exercise 2")
	/*
		Print every rune code point of the uppercase alphabet three times.
		Your output should look like this:
		65
			U+0041 'A'
			U+0041 'A'
			U+0041 'A'
		66
			U+0042 'B'
			U+0042 'B'
			U+0042 'B'
		 … through the rest of the alphabet characters
	*/
	for i := 65; i < 123; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#x\t%U\t%q\n", i, i, i)
		}
	}
}
func section9ex3() {
	fmt.Println("\nSection 9 Exercise 3")
	/*
		Create a for loop using this syntax
			for condition { }
		Have it print out the years you have been alive.
	*/
	var age int = 28
	for i := 1; i <= age; i++ {
		if i == age {
			fmt.Println("my age:", i)
		}
	}
}
func section9ex4() {
	fmt.Println("\nSection 9 Exercise 4")
	/*
		Create a for loop using this syntax
			for { }
		Have it print out the years you have been alive.
	*/
	var age int = 28
	var i int = 1
	for {
		if i == age {
			fmt.Println("my age:", i)
			break
		}
		i++
	}
}
func section9ex5() {
	fmt.Println("\nSection 9 Exercise 5")
	/*
		Print out the remainder (modulus) which is found for each number
		between 10 and 100 when it is divided by 4.
	*/
	for i := 10; i <= 100; i++ {
		quotient := i / 4
		remainder := i % 4
		fmt.Printf("%d / 4 = %dr%d\n", i, quotient, remainder)
	}
}
func section9ex6() {
	fmt.Println("\nSection 9 Exercise 6")
	/*
		Create a program that shows the “if statement” in action.
	*/
	x := rand.Intn(3)
	y := rand.Intn(3)
	if x == y {
		fmt.Println("x == y")
	}
}
func section9ex7() {
	fmt.Println("\nSection 9 Exercise 7")
	/*
		Building on the previous hands-on exercise, create a program that
		uses “else if” and “else”.
	*/
	x := rand.Intn(3)
	y := rand.Intn(3)
	if x == y {
		fmt.Println("x == y")
	} else if x > y {
		fmt.Println("x > y")
	} else {
		fmt.Println("x < y")
	}
}
func section9ex8() {
	fmt.Println("\nSection 9 Exercise 8")
	/*
		Create a program that uses a switch statement with no switch
		expression specified.
	*/
	switch { // default true
	case 1 == 1:
		fmt.Println("always true")
		fallthrough
	default:
		fmt.Println("default from fall through")
	}
}
func section9ex9() {
	fmt.Println("\nSection 9 Exercise 9")
	/*
		Create a program that uses a switch statement with the switch
		expression specified as a variable of TYPE string with the IDENTIFIER
		“favSport”.
	*/
	var favSport string = "surfing"
	switch favSport {
	case "mountain biking", "hiking":
		fmt.Println("mountain sports")
	case "surfing", "paddling":
		fmt.Println("water sports")
	default:
		fmt.Println("does not like sport")
	}
}
