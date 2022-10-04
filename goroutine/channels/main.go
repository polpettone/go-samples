package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const LINE_LENGTH int = 15

type sample func()

func main() {

	run(simpleGo)

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
	channel <- "aa"

	k := <-channel
	fmt.Printf("%s\n", k)

	channel <- "b"
	k = <-channel

	fmt.Printf("%s\n", k)
}

func writeToFullChannel() {
	fmt.Println(getFunctionName())

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
	channel <- "a"
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

	for n := 0; n < 3; n++ {
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

func simpleGo() {
	go one()

	time.Sleep(2 * time.Second)
	fmt.Println("fertig")

}

func one() {
	fmt.Println("Hallo")
}

func waitGroups() {

	channel := make(chan string, 1)
	channel <- "a"
	close(channel)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for n := 0; n < 3; n++ {
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
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Foobar")

	}(wg)

	wg.Wait()
}

func doneChannel() {

	done := make(chan struct{})
	ch := make(chan string, 1)

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func(wg *sync.WaitGroup, done chan struct{}, ch chan string) {
		defer wg.Done()
		for {
			select {
			case k, ok := <-ch:
				if ok {
					fmt.Printf("Channel read: %s \n", k)
				} else {
					fmt.Println("Channel read not ok")
				}
			case _, ok := <-done:
				if ok {
					fmt.Printf("Done read: %s \n", "ok")
				} else {
					fmt.Println("Done read not ok")
					return
				}
			default:
				fmt.Println("default")

			}
		}
	}(wg, done, ch)

	go func(wg *sync.WaitGroup, done chan struct{}) {
		defer wg.Done()
		for {
			select {
			case _, ok := <-done:
				if ok {
					fmt.Printf("Done read: %s \n", "ok")
				} else {
					fmt.Println("Done read not ok")
					return
				}
			default:
				fmt.Println("default 2")

			}
		}
	}(wg, done)

	go func(wg *sync.WaitGroup, done chan struct{}) {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		close(done)
	}(wg, done)

	wg.Wait()
}

func getFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
