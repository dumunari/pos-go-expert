package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Teste Product", 100)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Price)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Teste Product", product.Name)
	assert.Equal(t, 100.0, product.Price)
}

func TestProducWithNameRequired(t *testing.T) {
	product, err := NewProduct("", 100.0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWithPriceRequired(t *testing.T) {
	product, err := NewProduct("Teste Product", 0.0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWithInvalidPrice(t *testing.T) {
	product, err := NewProduct("Teste Product", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
