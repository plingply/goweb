/*
 * @Author: 彭林
 * @Date: 2021-01-10 11:37:51
 * @LastEditTime: 2021-01-10 11:41:15
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/library/cache/main.go
 */

package cache

import (
	"goframe-web/library/response"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Get(r *ghttp.Request, keys interface{}) {
	key := gmd5.MustEncrypt(keys)
	if v, err := g.Redis().DoVar("GET", key); err == nil && v.Map() != nil {
		response.JsonExit(r, 0, "作文列表", v.Map())
	}
}

func Set(keys interface{}, value interface{}, time uint) {
	key := gmd5.MustEncrypt(keys)
	g.Redis().Do("SET", key, value)
	g.Redis().Do("EXPIRE", key, time)
}
