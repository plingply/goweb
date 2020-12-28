package class_member

import (
	"goframe-web/app/model"
	"goframe-web/app/service/class_member"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetClassMemeberList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")
	class_id := r.GetQueryUint("class_id")

	if result, total, err := class_member.GetClassMemeberList(school_id, campus_id, class_id, page, limit); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		var resp = make(map[string]interface{})
		resp["item"] = result
		resp["total"] = total
		resp["page"] = page
		resp["limit"] = limit
		response.JsonExit(r, 0, "ok", resp)
	}
}

func CreateClassMemeber(r *ghttp.Request) {

	var (
		data        *createClassMemberRequest
		createParam *class_member.CreateParam
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &createParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var classMember model.ClassMember
	if err := gconv.Struct(createParam, &classMember); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if result, err := class_member.CreateClassMember(&classMember); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}
