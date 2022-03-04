package repository

import (
	"errors"

	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/static"
	"gorm.io/gorm"
)

type DAO interface {
	FindTargetTrackingProduct() ([]model.Product, error)
	GetProductPrice(productId uint64) (*model.Price, error)
	UpdateProductPrice(productId uint64, newPrice string) (int64, error)
	LockProductToTrackPrice(productId uint64) (int64, error)
	UnlockProduct(productId uint64) (int64, error)
	UpdateProductStatusToFailed(productId uint64) (int64, error)
	GetAllUserFollowed(productId uint64) ([]model.User, error)
}

type dao struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) (DAO, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &dao{
		db: db,
	}, nil
}

func (d *dao) FindTargetTrackingProduct() ([]model.Product, error) {
	targetProduct := []model.Product{}
	if err := d.db.Where("status = ? AND delete_flg = ? AND updated_at <= (NOW() - INTERVAL 1 MINUTE)",
		static.PRODUCT_STATUS_TRACKED,
		static.DELETE_FLAG_FALSE).Find(&targetProduct).Error; err != nil {
		return nil, err
	}
	return targetProduct, nil
}

func (d *dao) GetProductPrice(productId uint64) (*model.Price, error) {
	price := model.Price{}
	if err := d.db.Where("product_id = ? AND delete_flg = ?",
		productId,
		static.DELETE_FLAG_FALSE).Order("updated_at DESC").First(&price).Error; err != nil {
		return nil, err
	}

	return &price, nil
}

func (d *dao) UpdateProductPrice(productId uint64, newPrice string) (int64, error) {
	price := model.Price{
		Price:     newPrice,
		ProductID: productId,
	}
	res := d.db.Save(&price)
	return res.RowsAffected, res.Error
}

func (d *dao) GetAllUserFollowed(productId uint64) ([]model.User, error) {
	return d.getAllUserFollowed(productId)
}

func (d *dao) LockProductToTrackPrice(productId uint64) (int64, error) {
	return d.updateProductStatus(productId, static.PRODUCT_STATUS_TRACKED, static.PRODUCT_STATUS_ON_TRACKING)
}

func (d *dao) UnlockProduct(productId uint64) (int64, error) {
	return d.updateProductStatus(productId, static.PRODUCT_STATUS_ON_TRACKING, static.PRODUCT_STATUS_TRACKED)
}

func (d *dao) UpdateProductStatusToFailed(productId uint64) (int64, error) {
	return d.updateProductStatus(productId, static.PRODUCT_STATUS_ON_TRACKING, static.PRODUCT_STATUS_TRACKING_FAILED)
}

func (d *dao) updateProductStatus(productId uint64, whereStatus, toStatus int) (int64, error) {
	result := d.db.Table("t_product").Where("id = ? AND status = ? AND delete_flg = ?",
		productId, whereStatus, static.DELETE_FLAG_FALSE).
		Update("status", toStatus)
	return result.RowsAffected, result.Error
}

func (d *dao) getAllUserFollowed(productId uint64) ([]model.User, error) {

	users := []model.User{}

	findSQL := `
		SELECT 
			A.id, full_name, email, gender, follow_type, lang
		FROM t_user as A
		JOIN t_follow as B
		ON A.id = B.user_id AND B.delete_flg = ?
		WHERE B.product_id = ? AND A.delete_flg = ?
	`

	if err := d.db.Raw(findSQL,
		static.DELETE_FLAG_FALSE,
		productId,
		static.DELETE_FLAG_FALSE).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
