/*
 * @Author: 彭林
 * @Date: 2020-12-25 11:46:02
 * @LastEditTime: 2020-12-25 12:35:13
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/course/course.go
 */
package course

import (
	"goframe-web/app/service/course"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func GetCourseList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")
	class_id := r.GetQueryUint("class_id")

	result, total, err := course.GetCourseList(school_id, campus_id, class_id, page, limit)

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
