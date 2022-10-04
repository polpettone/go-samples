package main

import (
	"fmt"
)

func main() {
	fmt.Println("main() called")
	a, b := add()
	fmt.Printf("%d %d", a, b)
}

func add() (int, int) {
	return 42, 23
}
