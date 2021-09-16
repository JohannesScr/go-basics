package main

import "fmt"

func main() {
	fmt.Println("### GROUPING DATA ###")
	// let's look at arrays
	array()
	// let's look at slices
	slices()
	// let's look at slicing a slice
	sliceASlice()
	// let's look at appending to a slice
	appendSlice()
	// let's delete from a slice
	deleteSlice()
	// let's make a slice
	makeSlice()
	// let's make a multi-dimensional slice
	multiDimensionalSlice()
	// let's look at the map type
	mapType()

	// exercises
	sec11ex1()
	sec11ex2()
	sec11ex3()
	sec11ex4()
	sec11ex5()
	sec11ex6()
	sec11ex7()
	sec11ex8()
	sec11ex9()
	sec11ex10()
}

func array() {
	fmt.Println("\n### Array")
	fmt.Println("with an array in go you specify it as")
	fmt.Println("\tvar identifier [size]type")
	fmt.Println("that means each of the elements are of the same type " +
		"and that the array has a fixed size")
	fmt.Println("var x [5]int")
	var x [5]int
	fmt.Println(x, "and has length len(x) =", len(x))
	fmt.Printf("%T\n", x)

	fmt.Println("arrays are 0 based index, therefore we can alter the " +
		"value by accessing it by its index")
	fmt.Println("x[3] = 42")
	x[3] = 42
	fmt.Println(x)
	fmt.Println("however, in Go rather use slices")
}

func slices() {
	fmt.Println("\n### Slices")
	fmt.Println("a slice is Go's method of making array's dynamic. since " +
		"Go is about ease of programming")
	fmt.Println("\tvar identifier = []type{...values}")
	fmt.Println("var x = []int{1, 2, 3, 4, 5}")
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println(x, "and has length len(x) =", len(x))
	fmt.Printf("%T\n", x)

	fmt.Println("printing a slice")
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])

	for index, value := range x { // short-hand is usually to use i and v
		fmt.Println(index, value)
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func sliceASlice() {
	fmt.Println("\n### Slicing Slices")
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println("x[i] = [0 1 2 3 4] as the index of each element")
	fmt.Println("x[:] =", x[:])
	fmt.Println("x[1:3] =", x[1:3],
		"from index 1 up to but not including index 3")
	fmt.Println("x[:3] =", x[:3])
	fmt.Println("x[2:] =", x[2:])
}

func appendSlice() {
	fmt.Println("\n### Append to a Slice")
	x := []int{4, 5, 7, 8, 12}
	fmt.Printf("the slice is created as \n"+
		"x := []int{4, 5, 7, 8, 12}\n"+
		"> %v\n", x)

	x = append(x, 77, 12, 13, 64)
	fmt.Printf("\nwe then append to the slice, adding as many elements "+
		"as we want to of the same type\n"+
		"x = append(x, 77, 12, 13, 64)\n"+
		"now the append function has returned a new slice of the same type "+
		"as the original slice, including the added elements\n"+
		"> %v\n", x)

	y := []int{43, 21, 8, 70, 9}
	x = append(x, y...)
	fmt.Printf("\nbut now how do we put two slices together, we have\n"+
		"y := []int{43, 21, 8, 70, 9}\n"+
		"to do this we need to 'unroll' the slice y with y...\n"+
		"x = append(x, y...)\n"+
		"> %v\n"+
		"so we append to x by unrolling y into it's individual elements\n", x)
}

func deleteSlice() {
	fmt.Println("\n### Delete to a Slice")
	x := []int{4, 5, 7, 8, 12}
	fmt.Printf("we define a slice\n"+
		"x := []int{4, 5, 7, 8, 12}\n"+
		"> %v\n", x)

	x = append(x[:2], x[3:]...)
	fmt.Printf("now we want to remove the 2nd index, so we create a \n"+
		"new slice by appending, but removing the 2nd index\n"+
		"x = append(x[:2], x[3:]...)\n"+
		"> %v\n", x)
}

func makeSlice() {
	fmt.Println("\n### Make to a Slice")
	x := make([]int, 0)
	fmt.Printf("we make an empty slice\n"+
		"x := make([]int, 0) // you must specify a initial size\n"+
		"> %v\n", x)
	fmt.Printf("the length is just the initial length of the slice.\n"+
		"we can append to make the slice larger.\n"+
		"the capacity of a sice is the default size appended or allocated "+
		"in the runtime."+
		"and is default to 0\ncap(x)\n> %v\nlen(x)\n> %v\n", cap(x), len(x))
	y := make([]int, 10, 12)
	fmt.Printf("y := make([]int, 10, 12)\n> %v\ncap(y) = %v and "+
		"len(y) = %v\n", y, cap(y), len(y))
	y = append(y, 11, 12, 13)
	fmt.Printf("y = append(y, 11, 12, 13)\n> %v\ncap(y) = %v and "+
		"len(y) = %v\nnow we see that when we went over the initial "+
		"capacity of 12,\ngo allocated another capacity of 12 so that it "+
		"would not need\nto destroy and reallocate memory to the underlying "+
		"array of the slice\n", y, cap(y), len(y))
	fmt.Println("loop through a slice")
	for index, value := range x {
		fmt.Printf("%v: %v\n", index, value)
	}
}

func multiDimensionalSlice() {
	fmt.Println("\n### A multi-dimensional Slice")
	dim1o1 := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(dim1o1)
	dim1o2 := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(dim1o2)

	dim2 := [][]string{dim1o1, dim1o2}
	fmt.Println(dim2)
}

func mapType() {
	fmt.Println("\n### A map type\nthe map is a key value store and " +
		"is an unordered list. Similar to dict is Python and object in JavaScript")
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	} // the composite literal or the type
	fmt.Printf(`m := map[string]int{
	"James": 32,
	"Miss Moneypenny": 27,
}  // the composite literal or the type
> %v`, m)
	fmt.Printf("Then to access some entry\n"+
		`m["James"]`+"\n> %v\n", m["James"])
	fmt.Printf("if the entry does not exist in the map then the \n"+
		"map automatically returns the zero value.\n"+
		`m["Barnabas"]`+"\n> %v\n", m["Barnabas"])
	v, ok := m["Barnabas"]
	fmt.Printf("so we need to check if the entry exists called \n"+
		"the comma ok idiom.\n"+`v, ok := m["Barnabas"]`+
		"\n%v, %v\n", v, ok)

	if v, ok := m["James"]; ok {
		fmt.Printf("This is the if print: %v\n", v)
	}

	fmt.Println("adding a new element to a map and range")
	m["name"] = 12
	nameVar := "new name"
	intVar := 32
	m[nameVar] = intVar
	fmt.Println(`
	m["name"] = 12
	nameVar := "new name"
	intVar := 32
	m[nameVar] = intVar
`)
	fmt.Println("noe to use a for loop over the map")
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}

	delete(m, "James")
	fmt.Printf("delete from a map\n"+`delete(m, "James")`+
		"\n> %v\n", m)
	fmt.Println("you can also delete keys that do not exist, you\n" +
		"can the comma ok idiom to check if it actually deleted something")
	delete(m, "frikkie")
	fmt.Println("delete(m, \"frikkie\")")
	fmt.Println(m)
}

func sec11ex1() {
	fmt.Println("\n section 11 exercise 1")
	/*
		Using a COMPOSITE LITERAL:
			create an ARRAY which holds 5 VALUES of TYPE int
			assign VALUES to each index position.
		Range over the array and print the values out.
		Print out the TYPE of the array
	*/
	x := [5]int{12, 74, 19, 14, 17}
	for i, v := range x {
		fmt.Printf("index %d of value %d in array of " +
			"type %T\n", i, v, v)
	}
}
func sec11ex2() {
	fmt.Println("\n section 11 exercise 2")
	/*
		Using a COMPOSITE LITERAL:
			create a SLICE of TYPE int
			assign 10 VALUES
		Range over the slice and print the values out.
		Print out the TYPE of the slice
	*/
	x := []int{1, 2, 3, 4, 9, 8, 7, 6, 3, 10}
	for i, v := range x {
		fmt.Printf("index %d of value %d in slice of " +
			"type %T\n", i, v, v)
	}
}
func sec11ex3() {
	fmt.Println("\n section 11 exercise 3")
	/*
		Using the code from the previous example, use SLICING to create the following new slices which are then printed:
			[42 43 44 45 46]
			[47 48 49 50 51]
			[44 45 46 47 48]
			[43 44 45 46 47]
	*/
	s1 := []int{42, 43, 44, 45, 46}
	s2 := []int{47, 48, 49, 50, 51}
	s3 := []int{44, 45, 46, 47, 48}
	s4 := []int{43, 44, 45, 46, 47}
	fmt.Println(s1, s2, s3, s4)
}
func sec11ex4() {
	fmt.Println("\n section 11 exercise 4")
	/*
		Follow these steps:
			start with this slice
				x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
			append to that slice this value
				52
			print out the slice
			in ONE STATEMENT append to that slice these values
				53
				54
				55
			print out the slice
			append to the slice this slice
				y := []int{56, 57, 58, 59, 60}
			print out the slice
	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
func sec11ex5() {
	fmt.Println("\n section 11 exercise 5")
	/*
		To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
			start with this slice
				x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
			use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
			[42, 43, 44, 48, 49, 50, 51]

	*/
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Println(y)
}
func sec11ex6() {
	fmt.Println("\n section 11 exercise 6")
	/*
		Create a slice to store the names of all of the states in the United
		States of America. What is the length of your slice? What is the
		capacity? Print out all of the values, along with their index
		position in the slice, without using the range clause.
		Here is a list of the states:
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`,
		` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`,
		` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`,
	*/
	x := make([]string, 50, 50)
	x = []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`,
		` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`,
		` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`,
		` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`,
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`,
		` Wisconsin`, ` Wyoming`,
	}
	fmt.Printf("length: %d and capacity: %d\n", len(x), cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Printf("%d: %v\n", i, x[i])
	}
}
func sec11ex7() {
	fmt.Println("\n section 11 exercise 7")
	/*
		Create a slice of a slice of string ([][]string). Store the following
		data in the multi-dimensional slice:
			"James", "Bond", "Shaken, not stirred"
			"Miss", "Moneypenny", "Helloooooo, James."
		Range over the records, then range over the data in each record.
	*/
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{x1, x2}
	fmt.Println(x)
}
func sec11ex8() {
	fmt.Println("\n section 11 exercise 8")
	/*
		Create a map with a key of TYPE string which is a person’s
		“last_first” name, and a value of TYPE []string which stores
		their favorite things. Store three records in your map.
		Print out all of the values, along with their index position
		in the slice.

		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
	*/
	x := map[string][]string{
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)
	for key, value := range x {
		fmt.Println(key)
		for i, v := range value {
			fmt.Printf("%d: %s\n", i, v)
		}
	}
}
func sec11ex9() {
	fmt.Println("\n section 11 exercise 9")
	/*
		Using the code from the previous example, add a record to your map.
		Now print the map out using the “range” loop
	*/
	x := map[string][]string{
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)
	x["hood_robin"] = []string{"uno", "dos"}
	for key, value := range x {
		fmt.Println(key)
		for i, v := range value {
			fmt.Printf("%d: %s\n", i, v)
		}
	}
}
func sec11ex10() {
	fmt.Println("\n section 11 exercise 10")
	/*
		Using the code from the previous example, delete a record from your
		map. Now print the map out using the “range” loop
	*/
	x := map[string][]string{
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(x)
	x["hood_robin"] = []string{"uno", "dos"}
	delete(x, "no_dr")
	for key, value := range x {
		fmt.Println(key)
		for i, v := range value {
			fmt.Printf("%d: %s\n", i, v)
		}
	}
}
