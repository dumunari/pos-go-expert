package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &RepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount must be greater than 0"))

	err := CalculateTaxAndSave(10000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.EqualError(t, err, "amount must be greater than 0")
	repository.AssertExpectations(t)

	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
