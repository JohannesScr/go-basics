package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"sort"
)

type user struct {
	FirstName string
	LastName  string
	Age       int
} // marshal encoding needs the fields to be TitleCase

type sysUser struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
	Depth     bool   `json:"Depth"`
}

func main() {
	fmt.Printf("### JSON ###\n")
	// lets marshal data to json
	marshalJSON()
	// let's unmarshal json to data
	unmarshalJSON()
	// let's look at the writer interface
	writerInterface()
	// let's sort
	sortInGo()
	// let's implement a custom sort
	customSort()
	// let's look at bcrypt
	bcryptInGo()
	bcryptCheck()

	sec19ex1()
	sec19ex2()
	sec19ex3()
	sec19ex4()
	sec19ex5()
}

func marshalJSON() {
	fmt.Printf("\n### Marshal JSON ###\n")
	fmt.Printf("it is important to note that truct fields need to be" +
		"\nin TitleCase for it to be marshalled.\n\n")
	u1 := user{
		FirstName: "james",
		LastName:  "bond",
		Age:       32,
	}
	u2 := user{
		FirstName: "miss",
		LastName:  "moneypenny",
		Age:       27,
	}
	users := []user{u1, u2}
	fmt.Println(users)

	byteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))
}

func unmarshalJSON() {
	fmt.Printf("\n### Unmarshal JSON ###\n")
	jsonString := `[{"FirstName":"james","LastName":"bond","Age":32, "Depth": true},{"FirstName":"miss","LastName":"moneypenny","Age":27}]`
	jsonBytes := []byte(jsonString)
	fmt.Printf("%T\t%s\n", jsonString, jsonString)
	fmt.Printf("%T\t%v\n", jsonBytes, jsonBytes)

	var users []sysUser
	err := json.Unmarshal(jsonBytes, &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\t%v\n", users, users)
}

func writerInterface() {
	fmt.Printf("\n### Writer Interface ###\n")
	//fmt.Fprintln(os.Stdout, "Hello")

	f, err := ioutil.ReadFile("my-json.json")
	fmt.Printf("%T: %s == %v and %v errors\n", f, f, f, err)
	fmt.Println("now that we have read (from a Reader) the file as " +
		"bytes,\nwe can now unmarshal the bytes to json.")

	var users []user
	err = json.Unmarshal(f, &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}

func sortInGo() {
	fmt.Printf("\n### Sorting a Slice ###\n")
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", ".", "Moneypenny", "Dr. No", "m", "!"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("{%s: %d}", p.Name, p.Age)
}

// ByAge is a type: the slice of Person can also be used by the
// ByAge type and all the associated methods
type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func customSort() {
	fmt.Printf("\n### Custom Sort ###\n")
	fmt.Println("by adding the String method to a struct we are " +
		"defining\nthe string that is returned when the print functions" +
		"are\ncalled.")

	p1 := Person{Name: "James Bond", Age: 29}
	p2 := Person{Name: "Miss Moneypenny", Age: 25}
	p3 := Person{Age: 99}
	p4 := Person{Name: "M", Age: 34}
	p5 := Person{Name: "Dr. No", Age: 41}

	people := []Person{p1, p2, p3, p4, p5}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

func bcryptInGo() {
	fmt.Printf("\n### Go Bcrypt ###\n")
	s := "password123"
	cost := 4
	h, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("password: %s\n", s)
	fmt.Printf("hash: %s\n", h)
	fmt.Printf("hash bytes: %v\n", h)
}

func bcryptCheck() {
	fmt.Printf("\n### Go Bcrypt Check Password ###\n")
	h := "$2a$04$D/TSCVVu4S3zPd1iiH9eBObC/g99WvAww5FkbY50h.1JcHqy8agHG"
	s1 := "password123"
	s2 := "password1234"
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(s1))
	if err != nil {
		fmt.Println(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(h), []byte(s2))
	if err != nil {
		fmt.Println(err)
	}
}

type UserEx1 struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func sec19ex1() {
	fmt.Println("\n### Section 19 Exercise 1 ###")
	/*
		Starting with this code, marshal the []user to JSON. There is a
		little curve-ball in this hands-on exercise
			- remember to ask yourself what you need to do to EXPORT a
		value from a package.
	*/
	u1 := UserEx1{First: "James", Age: 32}
	u2 := UserEx1{First: "Moneypenny", Age: 27}
	u3 := UserEx1{First: "M", Age: 54}
	users := []UserEx1{u1, u2, u3}
	fmt.Println(users)
	userBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", string(userBytes))
}

type userEx2 struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func (u userEx2) String() string {
	return fmt.Sprintf("{First: %s Last: %s Age: %d Sayings: %v}",
		u.First, u.Last, u.Age, u.Sayings)
}
func sec19ex2() {
	fmt.Println("\n### Section 19 Exercise 2 ###")
	/*
		Starting with this code, unmarshal the JSON into a Go data structure.
	*/
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	var usersEx2 []userEx2
	err := json.Unmarshal([]byte(s), &usersEx2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersEx2)
}

type userEx3 struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func sec19ex3() {
	fmt.Println("\n### Section 19 Exercise 3 ###")
	/*
		Starting with this code, encode to JSON the []user sending the
		results to Stdout. Hint: you will need to use
		json.NewEncoder(os.Stdout).encode(v interface{})
	*/
	u1 := userEx3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := userEx3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := userEx3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []userEx3{u1, u2, u3}
	fmt.Println(users)
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
func sec19ex4() {
	fmt.Println("\n### Section 19 Exercise 4 ###")
	/*
		Starting with this code, sort the []int and []string for each person.
	*/
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type userEx5 struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func (u userEx5) sortSayings() []string {
	sort.Strings(u.Sayings)
	return u.Sayings
}

type ByAgeEx5 []userEx5

func (a ByAgeEx5) Len() int           { return len(a) }
func (a ByAgeEx5) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAgeEx5) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLastEx5 []userEx5

func (a ByLastEx5) Len() int           { return len(a) }
func (a ByLastEx5) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastEx5) Less(i, j int) bool { return a[i].Last < a[j].Last }

func sec19ex5() {
	fmt.Println("\n### Section 19 Exercise 5 ###")
	/*
		Starting with this code, sort the []user by
			- age
			- last
		Also sort each []string “Sayings” for each user
			- print everything in a way that is pleasant
	*/
	u1 := userEx5{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := userEx5{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := userEx5{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	usersEx5 := []userEx5{u1, u2, u3}
	fmt.Println("---------------------------")
	for _, u := range usersEx5 {
		fmt.Println(u.First, u.Last, u.Age)
		for i, s := range u.Sayings {
			fmt.Println("\t", i, s)
		}
	}
	fmt.Println("---------------------------")
	sort.Sort(ByAgeEx5(usersEx5))
	for _, u := range usersEx5 {
		u.sortSayings()
		fmt.Println(u.First, u.Last, u.Age)
		for i, s := range u.Sayings {
			fmt.Println("\t", i, s)
		}
	}
	fmt.Println("---------------------------")
	sort.Sort(ByLastEx5(usersEx5))
	for _, u := range usersEx5 {
		fmt.Println(u.First, u.Last, u.Age)
		for i, s := range u.Sayings {
			fmt.Println("\t", i, s)
		}
	}
}
