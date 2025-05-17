package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	income := 1000.0
	expectedTax := 10.0

	actualTax := CalculateTax(income)
	if actualTax != expectedTax {
		t.Errorf("Expected tax: %.2f, but got: %.2f", expectedTax, actualTax)
	}
}
