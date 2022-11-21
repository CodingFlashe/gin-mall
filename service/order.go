package service

import (
	"context"
	"fmt"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  int     `json:"order_num"  form:"order_num"` //订单号
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Type:      1, //默认未支付
		Money:     service.Money,
	}
	//检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID
	//订单号创建,自动生成的随机number+唯一标识的product id + 用户的 id
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum
	err = orderDao.CreateOrder(order)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderById(uint(orderId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orderDao := dao.NewOrderDao(ctx)

	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orderList, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orderList), uint(total))
}

func (service *OrderService) Delete(ctx context.Context, uId uint, oId string) serializer.Response {
	OrderId, _ := strconv.Atoi(oId)
	code := e.Success
	OrderDao := dao.NewOrderDao(ctx)
	err := OrderDao.DeleteOrderByOrderId(uint(OrderId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
