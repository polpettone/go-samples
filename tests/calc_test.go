package calc

import (
	"strconv"
	"testing"
)

func TestCalc(t *testing.T) {
	wanted := 3
	actual := Add(1, 2)
	if wanted != actual {
		t.Errorf("Wanted %d, got %d", wanted, actual)
	}
}

func TestCalcTableDriven(t *testing.T) {

	type Args struct {
		wanted int
		x      int
		y      int
	}

	tests := []Args{

		{
			x:      1,
			y:      2,
			wanted: 3,
		},

		{
			x:      1,
			y:      3,
			wanted: 4,
		},
	}

	for _, tt := range tests {

		t.Run(strconv.Itoa(tt.wanted), func(t *testing.T) {

			actual := Add(tt.x, tt.y)
			if actual != tt.wanted {
				t.Errorf("Wanted %d, got %d", tt.wanted, actual)
			}

		})
	}
}
