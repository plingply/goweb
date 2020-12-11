package middleware

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"goframe-web/library/jwt"
	"goframe-web/library/response"
)

// JWTAuth 中间件，检查token
func JWTAuth(r *ghttp.Request) {
	token := r.GetQueryString("token")
	if token == "" {
		token = r.GetFormString("token")
	}
	if token == "" {
		token = r.GetHeader("token")
	}
	fmt.Println("token:", token)
	if token == "" {
		response.JsonExit(r, 1, "身份已过期")
	}
	j := jwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == jwt.TokenExpired {
			response.JsonExit(r, 1, "授权已过期")
		}
		response.JsonExit(r, 1, err.Error())
	}
	r.SetCtxVar("username", claims.Username)
	r.Middleware.Next()
}