package main

type InMemoryRepo struct {
}

func (r InMemoryRepo) FindById(id int) (string, error) {
	return "Some Fake", nil
}

func (r InMemoryRepo) Find() {}
