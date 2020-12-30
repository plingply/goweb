package card

import (
	"goframe-web/app/model"
	"goframe-web/app/service/card"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetCardList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	result, total, err := card.GetCardList(school_id, campus_id, page, limit)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var resp = make(map[string]interface{})
	resp["item"] = result
	resp["total"] = total
	resp["page"] = page
	resp["limit"] = limit

	response.JsonExit(r, 0, "ok", resp)
}

func GetCardSimpleList(r *ghttp.Request) {

	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("school_id")

	result, err := card.GetCardSimpleList(school_id, campus_id)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func UpdateCard(r *ghttp.Request) {

	card_id := r.GetQueryUint("card_id")

	var data = make(map[string]interface{})

	if r.GetFormBool("card_name") {
		data["card_name"] = r.GetFormString("card_name")
	}

	if r.GetFormBool("remark") {
		data["remark"] = r.GetFormString("remark")
	}

	if r.GetFormBool("status") {
		data["status"] = r.GetFormString("status")
	}

	result, err := card.UpdateCard(card_id, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateCard(r *ghttp.Request) {

	var (
		data      *CardRequest
		cardParam *model.Card
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &cardParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	err := card.CreateCard(cardParam)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", nil)
}
