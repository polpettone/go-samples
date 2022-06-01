package main

import "fmt"

func main() {

	repo := InMemoryRepo{}

	service := Service{
		Repo: repo,
	}

	result, _ := service.DoSomething()
	fmt.Printf("Result %s", result)

}
