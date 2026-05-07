package repository

import (
	"github.com/app/gin-postgres-api/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id string) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id string) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Order("created_at desc").Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) FindByID(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepositoryImpl) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Product{}).Error
}
