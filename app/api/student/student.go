package student

import (
	"goframe-web/app/model"
	"goframe-web/app/service/student"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetStudentList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	result, total, err := student.GetStudentList(school_id, campus_id, page, limit)

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

func UpdateStudent(r *ghttp.Request) {

	student_id := r.GetQueryUint("student_id")

	var data = make(map[string]interface{})

	if r.GetForm("student_name") != nil {
		data["student_name"] = r.GetFormString("student_name")
	}

	if r.GetForm("sex") != nil {
		data["sex"] = r.GetFormUint("sex")
	}

	if r.GetForm("avatar") != nil {
		data["avatar"] = r.GetFormString("avatar")
	}

	if r.GetForm("address") != nil {
		data["address"] = r.GetFormString("address")
	}

	if r.GetForm("school_name") != nil {
		data["school_name"] = r.GetFormString("school_name")
	}

	if r.GetForm("birthday") != nil {
		data["birthday"] = r.GetFormUint("birthday")
	}

	if r.GetForm("remark") != nil {
		data["remark"] = r.GetFormString("remark")
	}

	if r.GetForm("status") != nil {
		data["status"] = r.GetFormUint("status")
	}

	result, err := student.UpdateStudent(student_id, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateStudent(r *ghttp.Request) {

	school_id := r.GetQueryUint("school_id")
	campus_id := r.GetQueryUint("campus_id")

	var (
		data         *StudentRequest
		studentParam *model.Student
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &studentParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	result, err := student.CreateStudent(school_id, campus_id, studentParam)

	if err != nil {
		response.JsonExit(r, 0, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}
