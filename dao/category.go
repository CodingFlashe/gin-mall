package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

// NewCategoryDao 把DB复制过来而已
func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 获取所有分类
func (dao *CategoryDao) ListCategory() (Category []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&Category).Error
	return
}
