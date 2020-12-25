/*
 * @Author: 彭林
 * @Date: 2020-12-25 14:52:03
 * @LastEditTime: 2020-12-25 16:14:11
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/peotry/peotry.go
 */
package peotry

import (
	"goframe-web/app/model"
	"goframe-web/app/service/peotry"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func CreatePeotry(r *ghttp.Request) {

	var params model.PeotryParams
	if err := r.Parse(&params); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var modelPeotry model.Peotry
	modelPeotry.Align = params.Align
	modelPeotry.Appreciate = params.Appreciate
	modelPeotry.Author = params.Author
	modelPeotry.AuthorMore = params.AuthorMore
	modelPeotry.DuYin = params.DuYin
	modelPeotry.OrgText = params.OrgText
	modelPeotry.PeotryId = params.PeotryId
	modelPeotry.Reason = params.Reason
	modelPeotry.Title = params.Title
	modelPeotry.Translation = params.Translation
	modelPeotry.Video = params.Video
	modelPeotry.Years = params.Years

	peotry.CreatePeotry(&modelPeotry, params.NoteList)

	response.JsonExit(r, 0, "创建成功", nil)
}

func GetPeotryLastId(r *ghttp.Request) {
	id := peotry.GetPeotryLastId()
	response.JsonExit(r, 0, "诗词id", id)
}

func GetPeotryList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")

	if result, total, err := peotry.GetPeotryList(page, limit); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {

		var resp = make(map[string]interface{})
		resp["item"] = result
		resp["total"] = total
		resp["page"] = page
		resp["limit"] = limit

		response.JsonExit(r, 0, "诗词列表", resp)
	}
}

func GetInfo(r *ghttp.Request) {
	peotry_id := r.GetFormUint("peotry_id")
	if data, err := peotry.GetInfo(peotry_id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "诗词详情", data)
	}
}
