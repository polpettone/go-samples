package main

import (
	"fmt"
	"testing"
)

type MockedRepo struct{}

func (r MockedRepo) FindById(id int) (string, error) {
	return "Mocked Result", nil
}

type MockedErrRepo struct{}

func (r MockedErrRepo) FindById(id int) (string, error) {
	return "", fmt.Errorf("Repo Error")
}

func TestDoSomething(t *testing.T) {

	type args struct {
		wanted    string
		Repo      Repo
		wantedErr error
	}

	tests := []args{

		{
			wanted:    "Some Fake",
			Repo:      InMemoryRepo{},
			wantedErr: nil,
		},

		{
			wanted:    "Mocked Result",
			Repo:      MockedRepo{},
			wantedErr: nil,
		},

		{
			wanted:    "",
			Repo:      MockedErrRepo{},
			wantedErr: fmt.Errorf("Repo Error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.wanted, func(t *testing.T) {

			service := Service{
				Repo: tt.Repo,
			}

			actual, err := service.DoSomething()

			if tt.wantedErr != nil {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("Wanted error: %s, got %s", tt.wantedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("Wanted no error got %s", err)
				}
			}

			if actual != tt.wanted {
				t.Errorf("Wanted %s, got %s", tt.wanted, actual)
			}

		})
	}

}
