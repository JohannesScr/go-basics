package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("### Pointers ###")
	// let's look at pointers and addresses
	basicAddressesAndPointers()
	whenToUsePointers()

	// pass a value
	outerValue()
	// pass a pointer
	outerPointer()
	// method sets
	methodSets()
	methodReceivers()

	sec17ex1()
	sec17ex2()
}

func basicAddressesAndPointers() {
	fmt.Println("### Basics of Pointers and Addresses ###")
	a := 42
	fmt.Println("a := 42")
	fmt.Printf("a will return the value %v, %T\n", a, a)
	fmt.Printf("&a will the address in memory %v, %T\n", &a, &a)
	fmt.Println("note that *int is a type. the type is a pointer " +
		"that points to an int")

	fmt.Println("\nnow everything is pass by value")
	fmt.Println("var b int = &a")
	fmt.Println("this will not work, since you cannot assign a type of" +
		"pointer int *int to a type of int.")
	fmt.Println("var b *int = &a")
	fmt.Println("this will work since the address of a (&a) is a\n" +
		"pointer to an int (*int) and we declare b as a point to an\n" +
		"int (*int)")

	b := &a
	fmt.Println("\nnow dereference an address")
	fmt.Println("b := &a")
	fmt.Printf("b = %v\n", b)
	fmt.Printf("*b = %v\n", *b)
	fmt.Println("\nnow we can assign a value to the address")
	*b = 43
	fmt.Printf("a: %v and &a: %v then b: %v and *b: %v\n", a, &a, b, *b)
	fmt.Println("when changing value of a pointer all the pointers will " +
		"change\nsince the value at the address is changed")
	fmt.Println("\n######")
}

func whenToUsePointers() {
	fmt.Println("### When to use Pointers ###")
	fmt.Println("- pointers gives a performance benefit when you want to " +
		"pass\na large chunk of data in your program. Then instead of " +
		"passing the\nlarge chunk data around all you need to do it to pass " +
		"the pointer\n(or address) to the data and the function will known " +
		"where to find it")
	fmt.Println("- when you need to change a value at a location.")
	fmt.Println("\n######")
}

func outerValue() {
	fmt.Println("# Pass Value")
	fmt.Println("identifier\taddress\t\tvalue")
	x := 0
	fmt.Printf("x\t\t%v\t\t%v\n", &x, x)
	passValue(x)
	fmt.Printf("x\t\t%v\t\t%v\n", &x, x)
}

func passValue(y int) {
	fmt.Printf("y\t\t%v\t\t%v\n", &y, y)
	y = 43
	fmt.Printf("y\t\t%v\t\t%v\n", &y, y)
}

func outerPointer() {
	fmt.Println("# Pass Pointer")
	fmt.Println("identifier\taddress\t\tvalue")
	x := 0
	fmt.Printf("x\t\t%v\t\t%v\n", &x, x)
	passPointer(&x)
	fmt.Printf("x\t\t%v\t\t%v\n", &x, x)
}

func passPointer(y *int) {
	fmt.Printf("y\t\t%v\t\t%v\n", y, *y)
	*y = 43
	fmt.Printf("y\t\t%v\t\t%v\n", y, *y)
}

func methodSets() {
	fmt.Println("\n### Pass Pointer ###")
	fmt.Printf("methods set determine what methods attach to a " +
		"type T\nIt is exactly what the name describes, a set of all the\n" +
		"methods attached to a type.\n\n")
	fmt.Println("Receivers\tValues")
	fmt.Println("(t T)\t\tT and *T")
	fmt.Println("(t *T)\t\tonly *T")

	fmt.Println("\nnote you cannot mix the pointer and non-pointer " +
		"\nreceiver types when mixed in a general interface")
	fmt.Println("\n######")
}

type circ struct {
	radius float64
}
type basicShape interface {
	area() float64
}

type basicShape2 interface {
	circumference() float64
}

func (ci circ) area() float64 {
	fmt.Println("area is a non-pointer receiver")
	return math.Pi * math.Pow(ci.radius, 2)
}
func (ci *circ) circumference() float64 {
	fmt.Println("circumference is a pointer receiver")
	return 2 * math.Pi * ci.radius
}
func infoArea(s basicShape) {
	fmt.Printf("%v has area of %v\n", s, s.area())
}
func infoCircumference(s basicShape2) {
	fmt.Printf("%v has circumference of %v\n", s, s.circumference())
}

func methodReceivers() {
	fmt.Println("\n### Method Receivers ###")
	c := circ{
		radius: 12.3,
	}
	fmt.Print(c.area(), c.circumference(), "\n\n")

	fmt.Printf("%T is %v\n", &c, &c)
	infoArea(c)
	infoArea(&c)
	infoCircumference(&c)
	fmt.Printf("%T is %v\n", &c, &c)
	fmt.Println("\n######")
}

func sec17ex1() {
	fmt.Println("\n### Section 17 Exercise 1 ###")
	/*
		Create a value and assign it to a variable.
		Print the address of that value.
	*/
	var x string
	x = "hi bob"
	fmt.Printf("%v is at address %v\n", x, &x)

	var y *string
	y = &x
	*y = "hi alice"
	fmt.Printf("%v is at address %v\n", *y, y)
	fmt.Printf("%v is at address %v\n", x, &x)

	// note you cannot declare a pointer and assign it directly.
	// you first need to assign the pointer an address before you can
	// change the value at the address
	var z *int
	b := 34
	z = &b
	fmt.Printf("%v is at address %v\n", z, z)
}

type humanoid struct {
	first string
}
func changeMe1(h *humanoid, name string) {
	h.first = name
}
func changeMe2(h *humanoid, name string) {
	(*h).first = name
}

func sec17ex2() {
	fmt.Println("\n### Section 17 Exercise 2 ###")
	/*
		create a person struct
		create a func called “changeMe” with a *person as a parameter
			- change the value stored at the *person address
			- important
				- to dereference a struct, use (*value).field
					- p1.first
					- (*p1).first
				- “As an exception, if the type of x is a named pointer type
				and (*x).f is a valid selector expression denoting a field
				(but not a method), x.f is shorthand for (*x).f.”
					- https://golang.org/ref/spec#Selectors
		create a value of type person
			- print out the value
		in func main
			- call “changeMe”
		in func main
			- print out the value
	*/
	var h humanoid
	h = humanoid{
		first: "bob",
	}
	fmt.Println(h)
	changeMe1(&h, "james")
	fmt.Println(h)
	changeMe1(&h, "freddy")
	fmt.Println(h)
}

