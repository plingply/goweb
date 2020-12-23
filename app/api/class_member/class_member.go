package class_member

import (
	"goframe-web/app/service/class_member"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
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
