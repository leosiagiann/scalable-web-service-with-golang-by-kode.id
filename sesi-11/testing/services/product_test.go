package services

import (
	"fmt"
	"testing"
	"testing/entity"
	"testing/repositories"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var ProductRepoMock = repositories.ProductRepositoryMock{Mock: mock.Mock{}}
var productServices = NewProductService(&ProductRepoMock)

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	id := "1"
	ProductRepoMock.Mock.On("FindById", id).Return(nil)

	product, err := productServices.GetOneProductById("1")

	require.Nil(t, product)
	require.NotNil(t, err)
	require.Equal(t, fmt.Sprintf("product with id %s not found", id), err.Error())
}

func TestProductServiceGetOneProductFound(t *testing.T) {
	id := "2"
	product := &entity.Product{
		Id:   id,
		Name: "product",
	}
	ProductRepoMock.Mock.On("FindById", id).Return(product)

	result, err := productServices.GetOneProductById("2")

	require.NotNil(t, result)
	require.Nil(t, err)
	require.Equal(t, result.Id, id)
	require.Equal(t, product.Name, result.Name)
}
