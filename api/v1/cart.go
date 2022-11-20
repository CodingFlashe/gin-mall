package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Cart api:", err)
	}
}

func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //这里的id是收藏夹的id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Cart api:", err)
	}
}

func ListCart(c *gin.Context) {
	listCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listCartService); err == nil {
		res := listCartService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("show product api:", err)
	}
}

func UpdateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), claim.ID, c.Param("id")) //第二个id是userId,最后的id是地址id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Cart api:", err)
	}
}
