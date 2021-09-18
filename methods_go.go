package main

import "fmt"

type personM1 struct {
	firstName string
	lastName string
}

type secretAgentM1 struct {
	personM1
	codeName string
	licenseToKill bool
}

func (s secretAgentM1) speak() {
	fmt.Printf("my name is %s, %s %s\n", s.lastName, s.firstName, s.lastName)
}
func (p personM1) speak() {
	fmt.Printf("my name is %s %s\n", p.firstName, p.lastName)
}

type human interface{
	speak()
}

func drive(h human) {
	fmt.Println("I am human and I can drive", h)
}

func iAm(h human) {
	switch h.(type) {
	case personM1:
		fmt.Printf("I am a person, %s %s\n",
			h.(personM1).firstName, h.(personM1).lastName)
	case secretAgentM1:
		fmt.Printf("I am a secret agent, shhh, %s, %s %s\n",
			h.(secretAgentM1).lastName,
			h.(secretAgentM1).firstName,
			h.(secretAgentM1).lastName)
	}
}

func main() {
	fmt.Println("### Methods ###")
	sa1 := secretAgentM1{
		personM1: personM1{
			firstName: "james",
			lastName: "bond",
		},
		codeName: "007",
		licenseToKill: true,
	}
	sa2 := secretAgentM1{
		personM1: personM1{
			firstName: "miss",
			lastName: "moneypenny",
		},
		codeName: "001",
		licenseToKill: true,
	}
	fmt.Println("function syntax: func (r receiver) identifier " +
		"(parameters) (returns) { ...code }")
	fmt.Println("in this case with secret agent and person as struct, " +
		"we can add methods to the structs\n" +
		"by defining the struct as the receiver of the the function which " +
		"is called a method.")
	fmt.Println("\nOur Secret Agent:", sa1)
	sa1.speak()
	fmt.Println("\nOur Secret Agent:", sa2)
	sa2.speak()

	fmt.Println("\n ### Interfaces ###")
	fmt.Println("now we have created a human of type interface and " +
		"assigned the method speak.\nthis now means that value that has a " +
		"speak method will also be of type human")
	fmt.Printf("%T\n", sa1)

	p1 := personM1{
		firstName: "Dr",
		lastName: "Yes",
	}
	drive(p1)
	drive(sa1)

	fmt.Println("\n### Interfaces and Polymorphism ###")
	fmt.Println("\nthis is known as polymorphism, where a person and a\n" +
		"secret agent are both also of type human, because they share an\n" +
		"interface drive")
	fmt.Println("\nnow we implement a switch statement, because even " +
		"though person\nand secretAgent share type human it does not mean\n" +
		"that both have the same fields")
	fmt.Println("\nwith the identifier.(type).field is a type of assertion")
	iAm(p1)
	iAm(sa1)
}