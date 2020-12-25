/*
 * @Author: 彭林
 * @Date: 2020-12-25 14:52:03
 * @LastEditTime: 2020-12-25 15:49:00
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
