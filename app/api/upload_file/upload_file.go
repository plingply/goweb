/*
 * @Author: 彭林
 * @Date: 2020-12-23 18:56:44
 * @LastEditTime: 2020-12-23 19:05:45
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/upload_file/upload_file.go
 */
package upload_file

import (
	"goframe-web/app/service/upload_file"
	"goframe-web/library/response"

	"github.com/gogf/gf/net/ghttp"
)

func UploadFile(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	name, err := upload_file.OSSUpLoad(file)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", name)
}
