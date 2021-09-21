package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func setupLog() *os.File {
	fmt.Println("create out log output file")
	f, err := os.Create(".log.txt")
	if err != nil {
		fmt.Println("log file error:", err)
	}
	log.SetOutput(f)

	//defer func() {
	//	fmt.Println("I am running defer")
	//	if err := f.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	return f
}

func main() {
	f := setupLog()
	defer f.Close()


	fmt.Println("### Go Error Handling ###")
	understandingErrorHanding()
	checkingErrors()
	//scanErr()
	//openFileErr()
	//errorTypesAndLogs()
	newErrorExample()
}

func understandingErrorHanding() {
	fmt.Println("\n### Type Error ###")
	fmt.Println("the error type is just an interface in go defined as:\n\n" +
		"\ttype error interface {\n" +
		"\t\tError() string\n" +
		"\t}\n\n" +
		"so any struct with a method `Error() string` is of type error.")
	fmt.Println("\n######")
}

func checkingErrors() {
	fmt.Println("\n### Type Error ###")
	fmt.Println("almost (in most cases) do error checking!")

	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	fmt.Println("\n######")
}

func scanErr() {
	fmt.Println("\n### Scan Error ###")
	var a1, a2, a3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&a1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&a2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&a3)
	if err != nil {
		panic(err)
	}
	fmt.Println(a1, a2, a3)
	fmt.Println("\n######")
}

func openFileErr() {
	fmt.Println("\n### Open File Error ###")
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("wassup")
	n, err := io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	fmt.Println("\n######")
}

func errorTypesAndLogs() {
	fmt.Println("\n### Open File Error and Logs ###")
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err:", err)
		log.Println("err:", err)
		//log.Fatal("err:", err)
		//panic(err)
	}
	fmt.Println("\n######")
}

func newErrorExample() {
	fmt.Println("\n### New Custom Error Example ###")
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n######")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}