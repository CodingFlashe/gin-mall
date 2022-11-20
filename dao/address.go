package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

// NewAddressDao 把DB复制过来而已
func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByAid(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).Find(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUserId(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&addresses).Error
	return
}
func (dao *AddressDao) UpdateAddressByUserId(aId uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id=?", aId).Updates(&address).Error
}
func (dao *AddressDao) DeleteAddressByAddressId(aId uint, uId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", aId, uId).Delete(&model.Address{}).Error
}
