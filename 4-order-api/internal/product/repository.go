package product

import (
	"github.com/Vozhlak/golang-advanced/4-order-api/pkg/db"
	"gorm.io/gorm/clause"
)

type RepositoryProduct struct {
	Db *db.Db
}

func NewRepositoryProduct(db *db.Db) *RepositoryProduct {
	return &RepositoryProduct{
		Db: db,
	}
}

func (r *RepositoryProduct) Create(product *Product) (*Product, error) {
	if result := r.Db.DB.Create(&product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *RepositoryProduct) GetProductById(id uint) (*Product, error) {
	var product Product
	if result := r.Db.DB.First(&product, id); result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *RepositoryProduct) Update(product *Product) (*Product, error) {
	if result := r.Db.DB.Clauses(clause.Returning{}).Updates(product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *RepositoryProduct) Delete(id uint) error {
	if result := r.Db.DB.Delete(&Product{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}
