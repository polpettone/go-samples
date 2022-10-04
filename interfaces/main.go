package main

import "fmt"

func main() {

	inMemoryRepo := InMemoryRepo{}

	service := Service{
		Repo: inMemoryRepo,
	}

	result, _ := service.DoSomething()
	fmt.Printf("Result %s", result)

}
