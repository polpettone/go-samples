package main

import "fmt"

func init() {
	fmt.Println("init() called")
}

func main() {
	fmt.Println("main() called")
}

func Never() {
	fmt.Println("never() called")
}
