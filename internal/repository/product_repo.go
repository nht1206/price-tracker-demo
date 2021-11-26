package repository

import (
	"fmt"

	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/static"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindTargetTrackingProduct() ([]model.Product, error)
	GetProductPrice(productId uint64) (*model.Price, error)
	UpdateProductPrice(productId uint64, newPrice string) (int64, error)
	LockProductToTrackPrice(productId uint64) (int64, error)
	UnlockProduct(productId uint64) (int64, error)
	UpdateProductStatusToFailed(productId uint64) (int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) (ProductRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	return &productRepository{
		db: db,
	}, nil
}

func (r *productRepository) FindTargetTrackingProduct() ([]model.Product, error) {
	targetProduct := []model.Product{}
	if err := r.db.Where("status = ? AND delete_flg = ? AND updated_at <= (NOW() - INTERVAL 1 MINUTE)",
		static.PRODUCT_STATUS_TRACKED,
		static.DELETE_FLAG_FALSE).Find(&targetProduct).Error; err != nil {
		return nil, err
	}
	return targetProduct, nil
}

func (r *productRepository) GetProductPrice(productId uint64) (*model.Price, error) {
	price := model.Price{}
	if err := r.db.Where("product_id = ? AND delete_flg = ?",
		productId,
		static.DELETE_FLAG_FALSE).Order("updated_at DESC").First(&price).Error; err != nil {
		return nil, err
	}

	return &price, nil
}

func (r *productRepository) UpdateProductPrice(productId uint64, newPrice string) (int64, error) {
	price := model.Price{
		Price:     newPrice,
		ProductID: productId,
	}
	res := r.db.Save(&price)
	return res.RowsAffected, res.Error
}

func (r *productRepository) LockProductToTrackPrice(productId uint64) (int64, error) {
	return r.updateProductStatus(productId, static.PRODUCT_STATUS_TRACKED, static.PRODUCT_STATUS_ON_TRACKING)
}

func (r *productRepository) UnlockProduct(productId uint64) (int64, error) {
	return r.updateProductStatus(productId, static.PRODUCT_STATUS_ON_TRACKING, static.PRODUCT_STATUS_TRACKED)
}

func (r *productRepository) UpdateProductStatusToFailed(productId uint64) (int64, error) {
	return r.updateProductStatus(productId, static.PRODUCT_STATUS_ON_TRACKING, static.PRODUCT_STATUS_TRACKING_FAILED)
}

func (r *productRepository) updateProductStatus(productId uint64, whereStatus, toStatus int) (int64, error) {
	result := r.db.Table("t_product").Where("id = ? AND status = ? AND delete_flg = ?",
		productId, whereStatus, static.DELETE_FLAG_FALSE).
		Update("status", toStatus)
	return result.RowsAffected, result.Error
}
