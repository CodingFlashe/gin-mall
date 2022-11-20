package v1

import (
	"gin-mall/pkg/util"
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := c.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Address api:", err)
	}
}

func ListAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Address api:", err)
	}
}

func DeleteAddress(c *gin.Context) {
	deleteAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //这里的id是收藏夹的id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Address api:", err)
	}
}

func ShowAddress(c *gin.Context) {
	showAddressService := service.AddressService{}
	if err := c.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("show product api:", err)
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(c.Request.Context(), claim.ID, c.Param("id")) //第二个id是userId,最后的id是地址id
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln("create Address api:", err)
	}
}
