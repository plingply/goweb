/*
 * @Author: 彭林
 * @Date: 2021-01-10 11:37:51
 * @LastEditTime: 2021-01-22 12:33:46
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/library/cache/main.go
 */

package cache

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Get(r *ghttp.Request, keys interface{}) (interface{}, error) {
	key := gmd5.MustEncrypt(keys)
	result, err := g.Redis().DoVar("GET", key)
	return result.Map(), err
}

func Set(keys interface{}, value interface{}, time uint) {
	key := gmd5.MustEncrypt(keys)
	g.Redis().Do("SET", key, value)
	if time > 0 {
		g.Redis().Do("EXPIRE", key, time)
	}
}
