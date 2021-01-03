package school

import (
	"goframe-web/app/model"
	"goframe-web/app/service/school"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func CreateSchool(r *ghttp.Request) {

	var data *model.SchoolParams
	var schoolParam *model.School
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := gconv.Struct(data, &schoolParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	campus_name := r.GetFormString("campus_name")
	userId := r.GetCtxVar("user_id").Uint()

	if err := school.CreateSchool(campus_name, userId, schoolParam); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "创建成功", nil)
	}
}
