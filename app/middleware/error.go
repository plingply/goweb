package middleware

import (
	"goframe-web/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log().Error("我发生了错误 中间件")
		//返回固定的友好信息
		r.Response.ClearBuffer()
		response.JsonExit(r, 1, err.Error())
	}
}
