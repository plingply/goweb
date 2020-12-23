package subject

import (
	"goframe-web/app/model"
	"goframe-web/app/service/subject"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetSubjectList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	result, total, err := subject.GetSubjectList(school_id, campus_id, page, limit)

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

func UpdateSubject(r *ghttp.Request) {

	subject_id := r.GetQueryUint("subject_id")

	var data = make(map[string]interface{})

	if r.GetFormBool("subject_name") {
		data["subject_name"] = r.GetFormString("subject_name")
	}

	if r.GetFormBool("remark") {
		data["remark"] = r.GetFormString("remark")
	}

	if r.GetFormBool("status") {
		data["status"] = r.GetFormString("status")
	}

	result, err := subject.UpdateSubject(subject_id, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateSubject(r *ghttp.Request) {

	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	var (
		data         *SubjectRequest
		subjectParam *model.Subject
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &subjectParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	result, err := subject.CreateSubject(school_id, campus_id, subjectParam)

	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}
