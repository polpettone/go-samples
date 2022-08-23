package main

import (
	"fmt"
	"os"
	"os/exec"
)

func setupShellSettings() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func inputKeyReceiver(keyChannel chan string) {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		i := string(b)
		keyChannel <- i
	}
}

func inputKeyHandler(keyChannel chan string) {

	for {
		select {

		case key := <-keyChannel:

			switch key {
			case "q":
				fmt.Printf("%s", "bye bye")
				os.Exit(0)
			case "r":
				fmt.Printf("%s", "reload \n")
			default:
				fmt.Printf("no binding for %s\n", key)
			}

		}
	}

}

func main() {
	fmt.Println("Start")
	setupShellSettings()

	keyChannel := make(chan string, 1)
	go inputKeyReceiver(keyChannel)
	go inputKeyHandler(keyChannel)

	select {}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
