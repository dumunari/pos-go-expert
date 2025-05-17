package tax

import (
	"testing"
)

// go test -coverprofile=coverage.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out
// go tool cover -html=coverage.out -o coverage.html
func TestCalculateTaxes(t *testing.T) {
	type calculateTaxes struct {
		income      float64
		expectedTax float64
	}

	taxes := []calculateTaxes{
		{income: 1000.0, expectedTax: 5.0},
		{income: 10000.0, expectedTax: 10.0},
		{income: 20000.0, expectedTax: 10.0},
		{income: 5000.0, expectedTax: 5.0},
		{income: 0.0, expectedTax: 0.0},
	}

	for _, tax := range taxes {
		actualTax := CalculateTax(tax.income)
		if actualTax != tax.expectedTax {
			t.Errorf("Expected tax: %.2f, but got: %.2f for income: %.2f", tax.expectedTax, actualTax, tax.income)
		}
	}
}
