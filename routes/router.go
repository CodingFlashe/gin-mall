package routes

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())                    //使用跨域的中间件
	r.StaticFS("/static", http.Dir("./static")) //加载静态文件
	v1 := r.Group("api/v1/")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		authed := v1.Group("/") // 需要登陆保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			//绑定邮箱
			authed.POST("user/sending-email", api.SendEmail)
			//验证邮箱
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)
		}
	}
	return r
}
