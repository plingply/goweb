package middleware

import (
	"fmt"
	"goframe-web/app/model"
	"goframe-web/library/jwt"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// JWTAuth 中间件，检查token 1
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
		response.JsonExit(r, 50008, "身份已过期")
	}

	j := jwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == jwt.TokenExpired {
			response.JsonExit(r, 50008, err.Error())
		}
		response.JsonExit(r, 50012, err.Error())
	}
	fmt.Println("claims:", claims)
	r.SetCtxVar("username", claims.Username)
	r.SetCtxVar("user_id", claims.Userid)

	// 验证系统token
	userId := r.GetCtxVar("user_id").Uint()
	var usertoken model.UserToken
	if bol := usertoken.Vtoken(userId, token); !bol {
		response.JsonExit(r, 1, "身份已过期")
	}

	r.Middleware.Next()
}
