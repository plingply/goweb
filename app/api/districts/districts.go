/*
 * @Author: 彭林
 * @Date: 2020-12-31 15:38:48
 * @LastEditTime: 2020-12-31 15:56:57
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/districts/districts.go
 */
package districts

import (
	"goframe-web/app/service/districts"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func GetDistrictsList(r *ghttp.Request) {
	parent_id := r.GetUint("parent_id")
	if result, err := districts.GetDistrictsList(parent_id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}
