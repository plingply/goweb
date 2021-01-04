package school_user

import (
	"fmt"
	"goframe-web/app/model"
	"goframe-web/app/service/school_user"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func GetSchoolList(r *ghttp.Request) {
	user_id := r.GetCtxVar("user_id").Uint()
	result := school_user.GetSchoolList(user_id)
	response.JsonExit(r, 0, "ok", result)
}

func GetTeacherList(r *ghttp.Request) {

	page := r.GetQueryUint("page")
	limit := r.GetQueryUint("limit")
	schoolId := r.GetQueryUint("school_id")
	campusId := r.GetQueryUint("campus_id")

	result, total, err := school_user.GetTeacherList(schoolId, campusId, page, limit)

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

func UpdateTeacher(r *ghttp.Request) {

	teacher_id := r.GetQueryUint("teacher_id")

	var data = make(map[string]interface{})

	if r.GetForm("campus_name") != nil {
		data["campus_name"] = r.GetFormString("campus_name")
	}

	if r.GetForm("address") != nil {
		data["address"] = r.GetFormString("address")
	}

	if r.GetForm("sex") != nil {
		data["sex"] = r.GetFormString("sex")
	}

	if r.GetForm("birthday") != nil {
		data["birthday"] = r.GetFormString("birthday")
	}

	if r.GetForm("entry_at") != nil {
		data["entry_at"] = r.GetFormString("entry_at")
	}

	result, err := school_user.UpdateTeacher(teacher_id, data)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}

func CreateTeacher(r *ghttp.Request) {

	var (
		data         *TeacherRequest
		teacherParam *model.SchoolUser
	)

	schoolId := r.GetQueryUint("school_id")

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &teacherParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	teacherParam.SchoolId = schoolId
	campusId := teacherParam.CampusId

	fmt.Println("teacherParam: &v", teacherParam)

	result, err := school_user.CreateTeacher(campusId, teacherParam)

	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok", result)
}
