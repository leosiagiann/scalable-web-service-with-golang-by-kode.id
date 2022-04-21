package repositories

import (
	"testing/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) FindById(id string) *entity.Product {
	args := r.Mock.Called(id)

	if args.Get(0) == nil {
		return nil
	}

	product := args.Get(0).(*entity.Product)
	return product
}
