package main

import "fmt"

func init() {
	fmt.Println("init() called")
}

func main() {
	defer shutdown()
	fmt.Println("main() called")
}

func shutdown() {
	fmt.Println("shutdown")
}

func Never() {
	fmt.Println("never() called")
}
