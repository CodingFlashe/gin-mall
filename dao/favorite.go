package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

// NewFavoriteDao 把DB复制过来而已
func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorite(uId uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", uId).Find(&resp).Error
	return
}

func (dao *FavoriteDao) FavoriteExistOrNot(pId uint, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id=? AND user_id=?", pId, uId).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}

func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&in).Error
}

func (dao *FavoriteDao) DeleteFavorite(uId uint, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).
		Where("id=? AND user_id=?", fId, uId).Delete(&model.Favorite{}).Error
}