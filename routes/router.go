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

		// 商品操作
		v1.GET("products", api.ListProduct)     //获取商品列表
		v1.GET("products/:id", api.ShowProduct) //展示商品信息
		v1.GET("imgs/:id", api.ListProductImg)  //获取商品图片地址
		v1.GET("categories", api.ListCategory)  //商品分类

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

			// 商品操作
			authed.POST("product", api.CreateProduct)  //创建商品
			authed.POST("products", api.SearchProduct) //搜索商品

			// 收藏夹操作
			authed.GET("favorites", api.ListFavorite)          //展示内容
			authed.POST("favorites", api.CreateFavorite)       //创建内容
			authed.DELETE("favorites/:id", api.DeleteFavorite) //删除内容

			//地址操作
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddress) //获取详细地址,这里是获取表中的第几个id
			authed.GET("addresses", api.ListAddress)     //获取地址列表,这里是获取某一个用户的所有id
			authed.PUT("addresses/:id", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			//订单操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts", api.ListCart) //获取某个订单
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)

		}
	}
	return r
}
