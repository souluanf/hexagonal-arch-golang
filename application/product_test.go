package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/souluanf/hexagonal-arch-golang/application"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.New(t).Nil(err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be 0 to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Price = 10

	product.Status = ""
	_, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Status = application.DISABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Name = ""
	_, err = product.IsValid()
	require.Equal(t, "Name: non zero value required", err.Error())

	product.Name = "Product 1"
	product.Status = "invalid status"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than 0", err.Error())
}

func TestProduct_Getters(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED

	require.Equal(t, product.Id, product.GetId())
	require.Equal(t, product.Name, product.GetName())
	require.Equal(t, product.Price, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}
