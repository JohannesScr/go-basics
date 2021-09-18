package main

import (
	"fmt"
	"math"
)

func main() {
	// let's look at an anonymous function
	anonymousFunction()
	// let's look at function expressions
	funcExpression()
	// let's return a function from a function
	s := returnAFunction(12)
	i := s()
	fmt.Println(i)
	fmt.Println(retFunc()())
	// let's look at callback functions
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(evenSum(sumInt, ii...))
	fmt.Println(oddSum(sumInt, ii...))
	// let's look at closures
	closures()
	a := incrementor()
	b := incrementor()
	a()
	a()
	a()
	b()
	b()
	// let's look at recursion
	fmt.Println("\n### Recursion ###")
	n := factorial(10)
	fmt.Println(n)

	// exercises
	sec15ex1()
	sec15ex2()
	sec15ex3()
	sec15ex4()
	sec15ex5()
	sec15ex6()
	sec15ex7()
	sec15ex8()
	sec15ex9()
}

func anonymousFunction() {
	fmt.Println("\n### Anonymous Function ###")
	fmt.Println("an anonymous function is executed immediately, " +
		"therefore\nthere is no need to give it an identifier but it does " +
		"require\nthe parentheses to execute the function")
	func(x int) {
		fmt.Printf("this is an anonymous function that takes "+
			"parameter: x = %v\n", x)
	}(12) // the parenthesis here execute the function immediately
	fmt.Printf("######\n")
}

func funcExpression() {
	fmt.Println("\n### Function Expression ###")
	fmt.Println("identifier := func(parameters) (return) { ...code }")
	f := func() {
		fmt.Println("this is my function expression")
	}
	f()
	fmt.Printf("######\n")
}

func returnAFunction(x int) func() int {
	fmt.Println("\n### Function Expression ###")
	fmt.Println("take note of the variable x int here")
	return func() int {
		return x + 12
	}
}

func retFunc() func() int {
	return func() int {
		return 14
	}
}

func sumInt(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenSum(f func(z ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if (v % 2) == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func oddSum(f func(z ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if (v % 2) == 1 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func closures() {
	fmt.Println("\n### Function Expression ###")
	fmt.Println("we want to enclose a variable to within a scope")

	var x int
	fmt.Printf("here x is %d\n", x)
	x++
	fmt.Printf("here x is %d\n", x)
	{
		var x int
		fmt.Printf("here x is %d\n", x)
		x++
		fmt.Printf("here x is %d\n", x)
	}
	fmt.Printf("######\n")
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		fmt.Println("incrementor", x)
		return x
	}
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func sec15ex1() {
	fmt.Println("\n### Section 15 Exercise 1")
	/*
		create a func with the identifier foo that
			- takes in a variadic parameter of type int
			- pass in a value of type []int into your func (unfurl the []int)
			- returns the sum of all values of type int passed in
		create a func with the identifier bar that
			- takes in a parameter of type []int
			- returns the sum of all values of type int passed in
	*/
	xi := []int{1, 2, 2, 4, 5}
	fooEx := func(x ...int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}
	fmt.Println(fooEx(xi...))
	barEx := func(x []int) int {
		total := 0
		for _, v := range x {
			total += v
		}
		return total
	}
	fmt.Println(barEx(xi))

}
func sec15ex2() {
	fmt.Println("\n### Section 15 Exercise 2")
	/*
		Use the “defer” keyword to show that a deferred func runs after
		the func containing it exits.
	*/
	p1 := func() {
		fmt.Println("p1")
	}
	p2 := func() {
		fmt.Println("p2")
	}

	defer p1()
	p2()
}

type person1 struct {
	first string
	last  string
	age   int
}

func (p person1) speak() {
	fmt.Printf("Hi there, my name is %v %v, I am %d years old\n",
		p.first, p.last, p.age)
}
func sec15ex3() {
	fmt.Println("\n### Section 15 Exercise 3")
	/*
		Create a user defined struct with
			- the identifier “person”
			- the fields:
				- first
				- last
				- age
		attach a method to type person with
			- the identifier “speak”
			- the method should have the person say their name and age
		create a value of type person
		call the method from the value of type person
	*/
	p1 := person1{
		first: "james",
		last:  "bond",
		age:   12,
	}
	p1.speak()
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%v has the area of %f\n", s, s.area())
}
func sec15ex4() {
	fmt.Println("\n### Section 15 Exercise 4")
	/*
		create a type SQUARE
		create a type CIRCLE
		attach a method to each that calculates AREA and returns it
			-  circle area= π r 2
			- square area = L * W
		create a type SHAPE that defines an interface as anything that has the AREA method
		create a func INFO which takes type shape and then prints the area
		create a value of type square
		create a value of type circle
		use func info to print the area of square
		use func info to print the area of circle
	*/
	c := circle{
		radius: 2.5,
	}
	r := rectangle{
		width:  2,
		length: 3,
	}
	info(c)
	info(r)
}

func sec15ex5() {
	fmt.Println("\n### Section 15 Exercise 5")
	/*
		Build and use an anonymous func
	*/
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	func(x ...int) {
		total := 0
		for _, v := range x {
			if (v % 3) == 0 {
				total += v
			}
		}
		fmt.Println(total)
	}(y...)
}
func sec15ex6() {
	fmt.Println("\n### Section 15 Exercise 6")
	/*
		Assign a func to a variable, then call that func
	*/
	f := func(x string) string {
		return x
	}
	fmt.Println(f("this is my func variable"))
}

func counter() func() int {
	x := 3
	return func() int {
		x++
		return x
	}
}
func sec15ex7() {
	fmt.Println("\n### Section 15 Exercise 7")
	/*
		Create a func which returns a func
		assign the returned func to a variable
		call the returned func
	*/
	myInc := counter()
	fmt.Println(myInc())
	fmt.Println(myInc())
	fmt.Println(myInc())
}

func seriesMod(mod int, x []int) []int {
	series := make([]int, 0)
	for i := 0; i < len(x); i++ {
		if (x[i] % mod) == 0 {
			series = append(series, x[i])
		}
	}
	return series
}
func sumMod(f func(int, []int) []int, mod int, x []int) int {
	total := 0
	ser := f(mod, x)
	for _, v := range ser {
		total += v
	}
	return total
}
func sec15ex8() {
	fmt.Println("\n### Section 15 Exercise 8")
	/*
		A “callback” is when we pass a func into a func as an argument. For this exercise,
			- pass a func into a func as an argument
	*/
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sumMod(seriesMod, 3, y))
}
func sec15ex9() {
	fmt.Println("\n### Section 15 Exercise 9")
	/*
		Closure is when we have “enclosed” the scope of a variable in
		some code block. For this hands-on exercise, create a func
		which “encloses” the scope of a variable:
	*/
	x := 12
	fmt.Println(x)
	{
		x := 23
		fmt.Println(x)
	}
	fmt.Println(x)
}
