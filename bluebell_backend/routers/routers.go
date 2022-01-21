package routers

import (
	"bluebell_backend/controller"
	"bluebell_backend/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)，用来添加项目状态
	// gin.Default 默认添加 GinLogger和 GinRecovery 的中间件，gin.New()则空
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)
	// 登录才能访问的路由,需添加上JWT中间件
	v1.Use(middlewares.JWTAuthMiddleware())

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.POST("/post", controller.CreatePostHandler)
	v1.GET("/post/:id", controller.GetPostDetailHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
