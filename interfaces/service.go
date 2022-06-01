package main

type Service struct {
	Repo Repo
}

func (s Service) DoSomething() (string, error) {
	return s.Repo.FindById(1)
}
