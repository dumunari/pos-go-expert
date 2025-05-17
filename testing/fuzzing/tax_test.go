package tax

import (
	"testing"
)

// go test -fuzz=.
// go test -run=FuzzCalculateTax/47a9c54cbd29f7c7
// go test -fuzz=. -fuzztime=10s
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-4, 0, 1000.0, 5000.0, 10000.0, 20000.0}
	for _, income := range seed {
		f.Add(income)
	}
	f.Fuzz(func(t *testing.T, income float64) {
		result := CalculateTax(income)
		if income <= 0 && result != 0 {
			t.Errorf("Expected tax to be 0 for income <= 0, but got: %.2f", result)
		}
		if income > 20000 && result != 20 {
			t.Errorf("Expected tax to be 20 for income > 20000, but got: %.2f", result)
		}
	})
}
