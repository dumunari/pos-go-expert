package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	income := 10000.0
	expectedTax := 10.0

	actualTax, err := CalculateTax(income)
	assert.Nil(t, err)
	assert.Equal(t, expectedTax, actualTax, "Expected tax: %.2f, but got: %.2f", expectedTax, actualTax)

	income = 0
	expectedTax = 0.0
	actualTax, err = CalculateTax(income)
	assert.Error(t, err, "income must be greater than 0")
	assert.Equal(t, expectedTax, actualTax, "Expected tax: %.2f, but got: %.2f", expectedTax, actualTax)
	assert.Contains(t, err.Error(), "greater than 0")
}
