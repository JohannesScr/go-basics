package main

import (
	"fmt"
)

func main() {
	fmt.Println("### Channels ###")
	understandingChannels()
	channelBuffers()
	directionalChannels()
	channelAndRange()
	channelsAndSelect()
}

func understandingChannels() {
	fmt.Println("\n### Understanding Channels ###")
	fmt.Println("channels block.\n\n" +
		"\tch1 := make(chan int)\n" +
		"\tch1 <- 42\n" +
		"\tfmt.Println(<-ch1)\n\n" +
		"since it blocks when it tries to pass the value to the channel.\n" +
		"there is no other location trying to read from the channel,\n" +
		"and we are stuck\n\n" +
		"\tch1 := make(chan int)\n" +
		"\tgo func(){\n" +
		"\t\tch1 <- 42 // blocks until it can pass the value\n" +
		"\t}()\n" +
		"\tfmt.Println(\"from channel 1:\", <-ch1) // blocks until it receives a value\n\n" +
		"this works since the goroutine is a different routine passing the\n" +
		"value 42 back to this routine.")
	ch1 := make(chan int)
	go func() {
		ch1 <- 42
	}()
	fmt.Println("from channel 1:", <-ch1)
	fmt.Println("\n######")
}
func channelBuffers() {
	fmt.Println("\n### Understanding Channels Buffers ###")
	// has a buffer that allows one value to sit in it
	ch2 := make(chan int, 1)
	ch2 <- 42
	fmt.Println(<-ch2)

	ch2 <- 42
	// if I try add another it will go into deadlock since the buffer
	//ch2 <- 43
	fmt.Println(<-ch2)

	fmt.Println("\nnext if channel buffers\n\n" +
		"\t// has a buffer that allows one value to sit in it\n" +
		"\tch3 := make(chan int, 2)\n" +
		"\tch3 <- 42\n" +
		"\tfmt.Println(<-ch3)\n\n" +
		"\tch3 <- 42\n" +
		"\tch3 <- 43\n" +
		"\tfmt.Println(<-ch3)\n" +
		"\tfmt.Println(<-ch3)")
	// has a buffer that allows one value to sit in it
	ch3 := make(chan int, 2)
	ch3 <- 42
	fmt.Println(<-ch3)

	ch3 <- 42
	ch3 <- 43
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
	fmt.Println("\n######")
}

func directionalChannels() {
	fmt.Println("\n### Directional Channels ###")
	fmt.Println("you can specify that a channel passed to a function,\n" +
		"can only receive from the channel or send something to a channel.\n" +
		"this makes reading and understanding the code a little easier.\n\n" +
		"\tch := make(<- chan int, 2)  // this is a receive-only channel\n" +
		"\tch := make(chan <- int, 2)  // this is a send-only channel\n\n" +
		"we read from left to right")
	ch1 := make(chan int, 2)
	//ch2 := make(chan int, 2)

	// send
	go sendChannel(ch1)
	// receive
	receiveChannel(ch1)

	fmt.Println("\n######")
}

func sendChannel(c chan <- int) {
	// here the channel is only send
	c <- 12
}
func receiveChannel(c <- chan int) {
	// here the channel is only receive
	fmt.Println(<-c)
}

func channelAndRange() {
	fmt.Println("\n### Channels and Range ###")
	fmt.Println("with the range, if you do not close the channel, the\n" +
		"range will try to read from a channel and go into deadlock.\n" +
		"therefore you need to close the channel to prevent the channel\n" +
		"from going into deadlock.")
	ch := make(chan int)
	go sendChRange(ch)
	receiveChRange(ch)
	fmt.Println("\n######")
}

func sendChRange(c chan <- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)  // close the channel
}

func receiveChRange(c <- chan int) {

	for v := range c {
		fmt.Println("Print from Channel and Range:", v)
	}
}

func channelsAndSelect() {
	fmt.Println("\n### Channels and Select ###")
	fmt.Println("is used to pull values off of channels.")
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go sendChSelect(even, odd, quit)

	// receive
	receiveChSelect(even, odd, quit)

	fmt.Println("\n######")
}

func sendChSelect(e, o, q chan <- int) {
	for i := 0; i < 10; i++ {
		if (i % 2) == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receiveChSelect(e, o, q <- chan int) {
	for {
		select {
		case ev := <- e:
			fmt.Println("even channel from select:", ev)
		case ov := <- o:
			fmt.Println("odd  channel from select:", ov)
		case qv := <- q:
			fmt.Println("quit channel from select:", qv)
			return
		}
	}
}