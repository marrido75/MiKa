package service

import (
	"mika/internal/database"
	"mika/internal/model"
)

func GetProducts(category string) ([]model.Product, int64) {
	var products []model.Product
	var total int64
	query := database.DB.Where("status = ?", "active")
	if category != "" {
		query = query.Where("category = ?", category)
	}
	query.Find(&products).Count(&total)
	return products, total
}

func GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func CreateProduct(p *model.Product) error {
	return database.DB.Create(p).Error
}

func UpdateProduct(p *model.Product) error {
	return database.DB.Save(p).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&model.Product{}, id).Error
}
