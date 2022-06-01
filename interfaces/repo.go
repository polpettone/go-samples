package main

type Repo interface {
	FindById(id int) (string, error)
}
