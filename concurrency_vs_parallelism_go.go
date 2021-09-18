package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func timer(f func()) {
	start := time.Now()
	//... operation that takes 20 milliseconds ...
	f()
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func main() {
	fmt.Println("### Concurrency vs Parallelism ###")
	fmt.Println("concurrency is a design pattern. this is where the " +
		"code has\nbeen written to be able to be run in parallel if the " +
		"CPU\nallows it to run in parallel. the amount of instructions in " +
		"parallel\nis determined by the CPU.")

	fmt.Print("\n\"do not communicate by sharing memory; instead, share" +
		"\nmemory by communicating\"\n\n")

	fmt.Println("FROM\t\truntime package")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	// let's look at wait groups
	waitGroup()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	// let's look at a basic channel
	basicChannel()
	// let's look at the race condition
	raceCondition()
	// let's look at mutex
	mutex()
	// let's look at atomic
	atomicInGo()

	sec21ex1()
	sec21ex2()
	sec21ex3()
	sec21ex4()
	sec21ex5()
	sec21ex6()
}

// loopUno is a function that prints out from 0 to 10
func loopUno() {
	for i := 0; i < 5; i++ {
		fmt.Println("uno", i)
	}
	wg.Done()
}

// loopDos is a function that prints out from 0 to 15
func loopDos() {
	for i := 0; i < 6; i++ {
		fmt.Println("dos", i)
	}
}

func waitGroup() {
	fmt.Println("\n### WaitGroup ###")
	fmt.Println("a wait group allows you to know when to wait for a " +
		"\ngoroutine to run")
	wg.Add(1)
	go loopUno()
	loopDos()
	wg.Wait()
	fmt.Printf("\n######\n")
}

// basicChannel creates a channel to return the value from the function in a
// goroutine
func basicChannel() {
	doSomething := func(x int) int {
		return x * 5
	}

	fmt.Println("\n### basicChannel ###")
	ch := make(chan int)
	go func() {
		fmt.Println("running as a go routine")
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
	fmt.Printf("\n######\n")
}

func raceCondition() {
	fmt.Println("\n### The Race Condition ###")
	fmt.Println("the race condition is where a variable address is " +
		"shared\nbetween two processes in parallel. they both read the " +
		"value of 0\nthen they both increment the counter with one. the" +
		"expected\nresult should have been two, since the value was " +
		"incremented\ntwice, however due to the lag between read and write" +
		"they\nread the same value and then write the same value.")

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)// denotes how many to wait for
	counter := 0
	fmt.Println("Before CPUs", runtime.NumCPU())
	fmt.Println("Before Goroutines", runtime.NumGoroutine())

	for i:= 0; i < gs; i++ {
		go func(){
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

	fmt.Printf("\n######\n")
}

func mutex() {
	fmt.Println("\n### Mutex ###")
	fmt.Println("mutex locks the access to a variable while it is in use")

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(gs)  // denotes how many to wait for
	fmt.Println("Before CPUs", runtime.NumCPU())
	fmt.Println("Before Goroutines", runtime.NumGoroutine())

	for i:= 0; i < gs; i++ {
		go func(){
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

	fmt.Printf("\n######\n")
}

func atomicInGo() {
	fmt.Println("\n### Atomic ###")
	fmt.Println("mutex locks the access to a variable while it is in use")

	var counter int64 = 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)  // denotes how many to wait for
	fmt.Println("Before CPUs", runtime.NumCPU())
	fmt.Println("Before Goroutines", runtime.NumGoroutine())

	for i:= 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

	fmt.Printf("\n######\n")
}

func sec21ex1() {
	fmt.Println("\n### Section 21 Exercise 1 ###")
	/*
		in addition to the main goroutine, launch two additional goroutines
			- each additional goroutine should print something out
			- use waitgroups to make sure each goroutine finishes before
			your program exists
	*/
	loops := 2

	var wg sync.WaitGroup
	wg.Add(loops)
	for i := 1; i <= loops; i++ {
		go func(x int) {
			fmt.Println("goroutine thing", x)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("about to exit")
}

type person21e2 struct {
	First string
}
func (p *person21e2) speak() {
	fmt.Printf("Say something %s, I'll make it up to you\n",
		p.First)
}
type human21e2 interface {
	speak()
}
func saySomething(h human21e2) {
	h.speak()
}
func sec21ex2() {
	fmt.Println("\n### Section 21 Exercise 2 ###")
	/*
		create a type person struct
		attach a method speak to type person using a pointer receiver
			- *person
		create a type human interface
			- to implicitly implement the interface, a human must have the speak method
		create func “saySomething”
			- have it take in a human as a parameter
			- have it call the speak method
		show the following in your code
			- you CAN pass a value of type *person into saySomething
			- you CANNOT pass a value of type person into saySomething
		here is a hint if you need some help
			- https://play.golang.org/p/FAwcQbNtMG
	*/

	p := person21e2{First: "jimmy"}
	//saySomething(p) // note the error here
	saySomething(&p)
}
func sec21ex3() {
	fmt.Println("\n### Section 21 Exercise 3 ###")
	/*
		Using goroutines, create an incrementer program
			- have a variable to hold the incrementer value
			- launch a bunch of goroutines
				- each goroutine should
					- read the incrementer value
						- store it in a new variable
					- yield the processor with runtime.Gosched()
					- increment the new variable
					- write the value in the new variable back to the
					incrementer variable
		use waitgroups to wait for all of your goroutines to finish
		the above will create a race condition.
		Prove that it is a race condition by using the -race flag
		if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
	*/

}
func sec21ex4() {
	fmt.Println("\n### Section 21 Exercise 4 ###")
	/*
		Fix the race condition you created in the previous exercise by using
		a mutex
			- it makes sense to remove runtime.Gosched()
	*/
}
func sec21ex5() {
	fmt.Println("\n### Section 21 Exercise 5 ###")
	/*
		Fix the race condition you created in exercise #4 by using package
		atomic
	*/
}
func sec21ex6() {
	fmt.Println("\n### Section 21 Exercise 6 ###")
	/*
		Create a program that prints out your OS and ARCH. Use the following
		commands to run it
			- go run
			- go build
			- go install
	*/
}
