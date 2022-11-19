package model

import (
	"gin-mall/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string //商品标题
	Info          string //商品信息
	ImgPath       string //商品展示图
	Price         string
	DiscountPrice string //折后价
	OnSale        bool   `gorm:"default:false"` //商品是否在售
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID)) //对商品进行加一操作
	// 增加商品的排名
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
