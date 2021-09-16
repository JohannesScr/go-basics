package main

import "fmt"

func main() {
	// let's look at functions
	basicSyntax()
	foo()
	bar("james")
	s := woo("moneypenny")
	fmt.Println(s)
	t1, t2, t3 := tang("james", "moneypenny")
	fmt.Println(t1, t2, t3)

	// let's look at variadic parameters
	variadicParameters()
	myIntPrint(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	myInterfacePrint(1, "2", 3, "4")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := sum(xi...)  // unrolling
	fmt.Println("the total is:", total)

	// let's look at defer
	functionDefer()
}

func basicSyntax() {
	fmt.Println("\n### Function Syntax ###")
	fmt.Println("> func (r receiver) identifier(parameters) (return(s)) { code... }")
	fmt.Println("when you define a function, you define parameters")
	fmt.Println("when you call a function, you pass in arguments")
	fmt.Printf("#######\n\n")
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("hello from woo,", s)
}

func tang(s string, t string) (string, string, bool) {
	return fmt.Sprint("hello from tang,", s),
		fmt.Sprint("olla el tang,", t),
		true
}

func variadicParameters() {
	fmt.Println("\n### Variadic Parameters ###")
	fmt.Println("variadic parameters means that we specify an arbitrary number of parameters")
	fmt.Println("> func (r receiver) identifier(parameters, ...<type>) (return(s)) { code... }")
	fmt.Println("the ...<type> means it is any number of parameters of the type <type>")
	fmt.Println("recall type interface is the underlying type for any type")
	fmt.Println("note that the type passed is a slice, hence we can use the index to access each value")
	fmt.Printf("#######\n\n")
}

func myIntPrint(x ...int) {
	fmt.Printf("%v\n%T\n", x, x)
}

func myInterfacePrint(x ...interface{}) {
	fmt.Printf("%v\n%T\n", x, x)
}

func sum(x ...int) int {
	s := 0
	for i, v := range x {
		s += v
		fmt.Printf("index %d: add %d to get a total of %d\n", i, v, s)
	}
	return s
}

func functionDefer() {}