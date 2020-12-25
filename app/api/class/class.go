package class

import (
	"goframe-web/app/model"
	"goframe-web/app/service/class"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetClassList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")
	user_id := r.GetCtxVar("user_id").Uint()

	result, total, err := class.GetClassList(school_id, campus_id, user_id, page, limit)

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

func GetClassSimpleList(r *ghttp.Request) {

	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("school_id")
	user_id := r.GetCtxVar("user_id").Uint()

	result, err := class.GetClassSimpleList(school_id, campus_id, user_id)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func UpdateClass(r *ghttp.Request) {

	class_id := r.GetQueryUint("class_id")

	var data = make(map[string]interface{})

	if r.GetFormBool("class_name") {
		data["class_name"] = r.GetFormString("class_name")
	}

	if r.GetFormBool("class_type") {
		data["class_type"] = r.GetFormString("class_type")
	}

	if r.GetFormBool("capacity") {
		data["capacity"] = r.GetFormUint("capacity")
	}

	if r.GetFormBool("remark") {
		data["remark"] = r.GetFormString("remark")
	}

	if r.GetFormBool("status") {
		data["status"] = r.GetFormString("status")
	}

	result, err := class.UpdateClass(class_id, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateClass(r *ghttp.Request) {

	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	var (
		data       *ClassRequest
		classParam *model.Classs
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &classParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	result, err := class.CreateClass(school_id, campus_id, classParam)

	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func GetClassInfo(r *ghttp.Request) {
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")
	class_id := r.GetQueryUint("class_id")

	if result, err := class.GetClassInfo(school_id, campus_id, class_id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "班级详情", result)
	}
}
