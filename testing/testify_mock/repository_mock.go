package tax

import (
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) SaveTax(tax float64) error {
	args := r.Called(tax)
	return args.Error(0)
}
