package router

import (
	"goframe-web/app/api/hello"
	"goframe-web/app/api/school_user"
	"goframe-web/app/api/user"
	"goframe-web/app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
		group.GET("/school/list", school_user.GetSchoolList)
	})
}
