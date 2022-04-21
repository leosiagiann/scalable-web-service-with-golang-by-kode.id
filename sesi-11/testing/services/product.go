package services

import (
	"fmt"
	"testing/entity"
	"testing/repositories"
)

type ProductService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetOneProductById(id string) (*entity.Product, error) {
	product := s.repo.FindById(id)

	if product == nil {
		return nil, fmt.Errorf("product with id %s not found", id)
	}

	return product, nil
}
