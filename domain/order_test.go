package domain

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	a := NewAccount("A")
	o := NewOrder(a)

	for i := 1; i < 10; i++ {
		o.AddProduct(NewProduct(fmt.Sprintf("Product #%d", i), float64(i)), 2)
	}

	assert.Equal(t, 45.0, o.TotalPrice())
}
