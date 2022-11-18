package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// NewNoticeDao 把DB复制过来而已
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
