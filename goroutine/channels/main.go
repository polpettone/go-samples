package main

import (
	"fmt"
	"runtime"
	"strings"
)

const LINE_LENGTH int = 15

type sample func()

func main() {

	run(closeChannel)

	fmt.Println(strings.Repeat("-", LINE_LENGTH))

}

func run(fn sample) {
	fmt.Println("Channel Samples")
	fmt.Println(strings.Repeat("-", LINE_LENGTH))
	fn()
	fmt.Println(strings.Repeat("-", LINE_LENGTH))

}

func writeAndRead() {
	fmt.Println(getFunctionName())

	channel := make(chan string, 1)
	channel <- "a"

	k := <-channel
	fmt.Printf("%s\n", k)

	channel <- "b"
	k = <-channel

	fmt.Printf("%s\n", k)
}

func writeToFullChannel(name string) {
	fmt.Println(name)

	channel := make(chan string, 1)

	channel <- "a"
	channel <- "b"

	k := <-channel
	fmt.Printf("%s\n", k)
}

func readFromEmptyChannel() {
	fmt.Println(getFunctionName())

	channel := make(chan string, 1)
	channel <- "a"

	k := <-channel
	fmt.Printf("%s\n", k)

	channel <- "b"
	k = <-channel

	fmt.Printf("%s\n", k)

	fmt.Println("Try get value from empty channel, will produce deadlock")
	k = <-channel
}

func selectStatement() {
	channel := make(chan string, 1)
	select {
	case k := <-channel:
		fmt.Printf(" ok %s\n", k)
	default:
		fmt.Printf("%s\n", "default")
	}
}

func closeChannel() {
	channel := make(chan string, 1)
	channel <- "a"
	close(channel)

	for n := 0; n < 2; n++ {
		select {
		case k, ok := <-channel:
			if ok {
				fmt.Printf("ok. value: %s \n", k)
			} else {
				fmt.Printf("!ok. value: %s \n", k)
			}
		default:
			fmt.Printf("%s\n", "default")
		}
	}
}

func getFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
