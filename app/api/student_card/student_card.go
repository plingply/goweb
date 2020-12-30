/*
 * @Author: 彭林
 * @Date: 2020-12-30 17:37:34
 * @LastEditTime: 2020-12-30 17:40:36
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/student_card/student_card.go
 */
package student_card

import (
	"goframe-web/app/service/student_card"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func ActivateCard(r *ghttp.Request) {
	var (
		data *student_card.StudentCardParams
	)

	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if id, err := student_card.ActivateCard(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "创建成功", id)
	}
}
