package main

import (
	"fmt"
)

func main() {
	// let's look at a struct
	structType()
	// let's embed a struct in another
	embeddedStruct()
	// let's look at an anonymous struct
	anonymousStruct()
	// section 13 exercises
	sec13ex1()
	sec13ex2()
	sec13ex3()
	sec13ex4()
}

/* SECTION 1 */
type person struct {
	first string
	last  string
}

func structType() {
	fmt.Println("\n### Struct Type")
	fmt.Println(`
we define a struct as:
type <struct-name> struct{
	<attribute-name> <type>
}
type person struct{
	first string
	last string
}
so we declare a new type called person which has a structure,
the first attribute is first of type string and the second attribute is 
called last of type string
`)
	p1 := person{
		first: "james",
		last:  "bond",
	}
	p2 := person{
		first: "miss",
		last:  "moneypenny",
	}
	fmt.Printf("%T: %v", p1, p1)
	fmt.Printf("%T: %v", p2, p2)
	fmt.Println("then we access the attributes using the dot notation")
	fmt.Printf("person.first = %s\n", p1.first)
}

type secretAgent struct {
	person        // implicit anonymous field that is person of type person
	licenseToKill bool
	first         string
}

func embeddedStruct() {
	fmt.Println("\n### Embedded Struct")

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		first:         "007",
		licenseToKill: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.licenseToKill)
	fmt.Println(sa1.first, sa1.person.first, sa1.person.last, sa1.licenseToKill)
	fmt.Println("It is important to note that go let's the outer\n" +
		"type struct inherit the attributes from the inner type struct.\n" +
		"If the outer type has similar attributes as the inner type then\n" +
		"needs to be accessed via the <outer type>.<inner type>." +
		"<field>.\nhowever best practice would be not to have " +
		"conflicting fields")
}

func anonymousStruct() {
	fmt.Println("\n### Anonymous Struct")
	type person struct {
		first string
		last  string
		age   int
	}
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("we see that p1 and p2 are the same. howver, p2 is \n" +
		"an anonymous struct, since the struct is declare with the " +
		"assignment.")
}

func sec13ex1() {
	fmt.Println("\n### Section 13 Exercise 1")
	/*
		Create your own type “person” which will have an underlying type
		of “struct” so that it can store the following data:
			first name
			last name
			favorite ice cream flavors
		Create two VALUES of TYPE person. Print out the values, ranging
		over the elements in the slice which stores the favorite flavors.
	*/
	type person struct {
		firstName string
		lastName string
		iceCreamFlavours []string
	}
	p1 := person{
		firstName: "james",
		lastName: "bond",
		iceCreamFlavours: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		firstName: "miss",
		lastName: "moneypenny",
		iceCreamFlavours: []string{"strawberry"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for _, flavour := range p1.iceCreamFlavours {
		fmt.Println(flavour)
	}
}
func sec13ex2() {
	fmt.Println("\n### Section 13 Exercise 2")
	/*
		Take the code from the previous exercise, then store the values of
		type person in a map with the key of last name. Access each value
		in the map. Print out the values, ranging over the slice.
	*/
	type person struct {
		firstName string
		lastName string
		iceCreamFlavours []string
	}
	p1 := person{
		firstName: "james",
		lastName: "bond",
		iceCreamFlavours: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		firstName: "miss",
		lastName: "moneypenny",
		iceCreamFlavours: []string{"strawberry"},
	}
	persons := make(map[string]person)
	persons[p1.lastName] = p1
	persons[p2.lastName] = p2
	fmt.Println(persons)
	for key, p := range persons {
		fmt.Println(key, ":", p)
		for _, flavour := range p.iceCreamFlavours {
			fmt.Printf("\t%v\n", flavour)
		}
	}

	fmt.Println("or as a composite literal")
	personsAgain := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(personsAgain)
	for key, p := range personsAgain {
		fmt.Println(key, ":", p)
		for _, flavour := range p.iceCreamFlavours {
			fmt.Printf("\t%v\n", flavour)
		}
	}
}
func sec13ex3() {
	fmt.Println("\n### Section 13 Exercise 3")
	/*
		Create a new type: vehicle.
			- The underlying type is a struct.
			- The fields:
				- doors
				- color
		Create two new types: truck & sedan.
			- The underlying type of each of these new types is a struct.
			- Embed the “vehicle” type in both truck & sedan.
			- Give truck the field “fourWheel” which will be set to bool.
			- Give sedan the field “luxury” which will be set to bool. solution
		Using the vehicle, truck, and sedan structs:
			- using a composite literal, create a value of type truck and
			assign values to the fields;
			- using a composite literal, create a value of type sedan and
			assign values to the fields.
		Print out each of these values.
		Print out a single field from each of these values.
	*/
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}
	fmt.Printf("%T\t%v\t\t%v\n", t1, t1, t1.doors)
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}
	fmt.Printf("%T\t%v\t%v\n", s1, s1, s1.doors)
}
func sec13ex4() {
	fmt.Println("\n### Section 13 Exercise 4")
	/*
		Create and use an anonymous struct.
	*/
	anon := struct{
		id int
		name string
		coins []int
		friends map[string]int
	}{
		id: 1,
		name: "james",
		coins: []int{1, 1, 2, 5, 5, 5, 5},
		friends: map[string]int{
			"Moneypenny": 555,
			"Q": 777,
			"M": 888,
		},
	}
	fmt.Println(anon)
}
