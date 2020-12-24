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

func GetZuowenList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")

	if result, total, err := zuowen.GetZuowenList(page, limit); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {

		var resp = make(map[string]interface{})
		resp["item"] = result
		resp["total"] = total
		resp["page"] = page
		resp["limit"] = limit

		response.JsonExit(r, 0, "作文列表", resp)
	}
}

func GetZuowenLastId(r *ghttp.Request) {

	if id, err := zuowen.GetZuowenLastId(); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "文章id", id)
	}
}

func GetInfo(r *ghttp.Request) {
	zuowen_id := r.GetFormUint("zuowen_id")
	if data, err := zuowen.GetInfo(zuowen_id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "作文详情", data)
	}
}
