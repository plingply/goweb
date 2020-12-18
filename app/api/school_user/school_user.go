package school_user

import (
	"goframe-web/app/service/school_user"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func GetSchoolList(r *ghttp.Request) {
	user_id := r.GetCtxVar("user_id").Uint()
	result := school_user.GetSchoolList(user_id)
	response.JsonExit(r, 0, "ok", result)
}
