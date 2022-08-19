package main

import "fmt"

type Element interface {
	GetID() string
}

type Item struct {
	ID string
}

func (i Item) GetID() string {
	return i.ID
}

type Player struct {
	ID string
}

func (i Player) GetID() string {
	return i.ID
}

func main() {

	var e *Element
	e = &Item{ID: "1"}

	fmt.Printf("%v", e)

	elems := []*Element{
		&Item{ID: "1"},
		&Player{ID: "1"},
		e,
	}

}
