package campus

import (
	"goframe-web/app/model"
	"goframe-web/app/service/campus"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetCampusList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	schoolId := r.GetQueryUint("school_id")
	user_id := r.GetCtxVar("user_id").Uint()

	result, total, err := campus.GetCampusList(schoolId, user_id, page, limit)

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

func GetCampusSimpleList(r *ghttp.Request) {

	schoolId := r.GetQueryUint("school_id")
	user_id := r.GetCtxVar("user_id").Uint()

	result, err := campus.GetCampusSimpleList(schoolId, user_id)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func UpdateCampus(r *ghttp.Request) {

	campusId := r.GetQueryUint("campus_id")

	var data = make(map[string]interface{})

	if r.GetFormBool("campus_name") {
		data["campus_name"] = r.GetFormString("campus_name")
	}

	if r.GetFormBool("address") {
		data["address"] = r.GetFormString("address")
	}

	if r.GetFormBool("province") {
		data["province"] = r.GetFormString("province")
	}

	if r.GetFormBool("city") {
		data["city"] = r.GetFormString("city")
	}

	if r.GetFormBool("area") {
		data["area"] = r.GetFormString("area")
	}

	result, err := campus.UpdateCampus(campusId, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateCampus(r *ghttp.Request) {

	schoolId := r.GetQueryUint("school_id")

	var (
		data        *CampusRequest
		campusParam *model.Campus
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &campusParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	result, err := campus.CreateCampus(schoolId, campusParam)

	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}
