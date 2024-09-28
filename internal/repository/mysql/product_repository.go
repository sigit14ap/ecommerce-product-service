package repository

import (
	"sync"

	"github.com/sigit14ap/product-service/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllByShopID(shopID uint64) ([]domain.Product, error)
	GetByIDAndShopID(id, shopID uint64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id, shopID uint64) error
}

type productRepository struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db:    db,
		mutex: sync.Mutex{},
	}
}

func (repository *productRepository) GetAllByShopID(shopID uint64) ([]domain.Product, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	var products []domain.Product
	if err := repository.db.Where("shop_id = ?", shopID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *productRepository) GetByIDAndShopID(id, shopID uint64) (*domain.Product, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	var product domain.Product
	if err := repository.db.Where("id = ? AND shop_id = ?", id, shopID).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *productRepository) Create(product *domain.Product) error {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	return repository.db.Create(product).Error
}

func (repository *productRepository) Update(product *domain.Product) error {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	return repository.db.Omit("CreatedAt").Save(product).Error
}

func (repository *productRepository) Delete(id, shopID uint64) error {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	return repository.db.Where("id = ? AND shop_id = ?", id, shopID).Delete(&domain.Product{}).Error
}
