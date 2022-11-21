package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Order api:", err)
	}
}

func DeleteOrder(c *gin.Context) {
	deleteOrderService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //这里的id是收藏夹的id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Order api:", err)
	}
}

func ListOrder(c *gin.Context) {
	listOrderService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("show product api:", err)
	}
}

func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id")) //第二个id是userId,最后的id是地址id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Order api:", err)
	}
}
