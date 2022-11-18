package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// NewCarouselDao 把DB复制过来而已
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// ListCarousel 根据id获取Carousel
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
