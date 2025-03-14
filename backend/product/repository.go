package product

import "gorm.io/gorm"

type RepositoryProduct interface {
	FindAll() ([]Product, error)
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindById(id int) (Product, error)
	Delete(product Product) (Product, error)
}

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db}
}

func (r *repositoryProduct) FindAll() ([]Product, error) {
	var product []Product

	err := r.db.Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) FindById(id int) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", id).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repositoryProduct) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repositoryProduct) Delete(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
