/*
 * @Author: 彭林
 * @Date: 2020-12-24 15:35:54
 * @LastEditTime: 2020-12-24 18:14:38
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/zuowen/zuowen.go
 */
package zuowen

import (
	"goframe-web/app/model"
	"goframe-web/app/service/zuowen"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func SaveZuowen(r *ghttp.Request) {

	var zuowenModel model.Zuowen

	if err := r.Parse(&zuowenModel); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if id := zuowen.CreateZuowen(&zuowenModel); id == 0 {
		response.JsonExit(r, 1, "同步失败")
	} else {
		response.JsonExit(r, 0, "同步成功", id)
	}
}
