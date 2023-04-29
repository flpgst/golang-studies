package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	product, err := NewProduct("Product1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product1", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
	assert.EqualError(t, ErrNameIsRequired, ErrNameIsRequired.Error())
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
	assert.EqualError(t, ErrPriceIsRequired, ErrPriceIsRequired.Error())
}

func TestProductWhenNameIsValid(t *testing.T) {
	product, err := NewProduct("Product1", -10)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
	assert.EqualError(t, ErrInvalidPrice, ErrInvalidPrice.Error())
}
