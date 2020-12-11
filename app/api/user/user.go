package user

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"goframe-web/app/service/user"
	"goframe-web/library/response"
)

func SignUp(r *ghttp.Request) {
	var (
		data        *SignUpRequest
		signUpParam *user.SignUpParam
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if result, err := user.SignUp(signUpParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 用户详情
func Info(r *ghttp.Request) {
	id := r.GetQueryUint("id")
	if result, err := user.GetUserInfo(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 修改用户信息
func UpdateInfo(r *ghttp.Request) {

	var reqMap = make(map[string]interface{})
	id := r.GetFormUint("id")

	if r.GetFormBool("nickname") {
		reqMap["nickname"] = r.GetFormString("nickname")
	}

	if r.GetFormBool("password") {
		reqMap["password"] = r.GetFormString("password")
	}

	if err := user.Update(id, reqMap); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", "修改成功")
	}
}

// 登录
func Login(r *ghttp.Request) {
	passport := r.GetFormString("passport")
	password := r.GetFormString("password")
	if result, err := user.Login(passport, password); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}

// 退出
func Signout(r *ghttp.Request) {

}
