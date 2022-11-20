package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

// NewCartDao 把DB复制过来而已
func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByCid(cId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", cId).Find(&cart).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&carts).Error
	return
}
func (dao *CartDao) UpdateCartByUserId(cId uint, cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Updates(&cart).Error
}
func (dao *CartDao) DeleteCartByCartId(cId uint, uId uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id=? AND user_id=?", cId, uId).
		Delete(&model.Cart{}).Error
}

func (dao *CartDao) UpdateCartNumById(cId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id=?", cId).
		Update("num", num).Error
}
