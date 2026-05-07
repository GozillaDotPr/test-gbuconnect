package service

import (
	"errors"

	"github.com/app/gin-postgres-api/internal/models"
	"github.com/app/gin-postgres-api/internal/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	Create(name, desc string, price int64, userID string) (*models.Product, error)
	Update(id string, name string, desc string, price int64) (*models.Product, error)
	Delete(id string) error
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductServiceImpl) GetByID(id string) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return product, err
}

func (s *ProductServiceImpl) Create(name, desc string, price int64, userID string) (*models.Product, error) {
	product := &models.Product{
		Name:   name,
		Desc:   desc,
		Price:  price,
		UserID: userID,
	}
	if err := s.repo.Create(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) Update(id string, name string, desc string, price int64) (*models.Product, error) {
	product, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	if name != "" {
		product.Name = name
	}
	if desc != "" {
		product.Desc = desc
	}
	if price != 0 {
		product.Price = price
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) Delete(id string) error {
	_, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("product not found")
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
