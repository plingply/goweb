package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goframe-web/app/api/hello"
	"goframe-web/app/api/user"
	"goframe-web/app/middleware"
)

func init() {
	s := g.Server()

	s.Use(middleware.MiddlewareErrorHandler)
	s.Use(middleware.CORS)
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.POST("/login", user.Login)
		group.POST("/signup", user.SignUp)
		group.Middleware(middleware.JWTAuth)
		group.GET("/user/info", user.Info)
		group.POST("/signout", user.Signout)
		group.POST("/user/update", user.UpdateInfo)
	})
}
