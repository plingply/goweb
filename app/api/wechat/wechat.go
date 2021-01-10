/*
 * @Author: 彭林
 * @Date: 2021-01-10 11:56:43
 * @LastEditTime: 2021-01-10 13:19:07
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/wechat/wechat.go
 */
package wechat

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Verification(r *ghttp.Request) {

	var data *messageEntity
	token := g.Cfg().GetString("wechat.token")

	if err := r.Parse(&data); err != nil {
		r.Response.WriteExit(false)
		r.Exit()
	}

	fmt.Println(data, token)

	// _ := data.CheckSignature(token)
	r.Response.WriteExit(true)
}

type messageEntity struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

func (m *messageEntity) CheckSignature(token string) bool {
	item := []string{token, m.Timestamp, m.Nonce}
	sort.Strings(item)
	itemByte := strings.Join(item, "")
	signature := fmt.Sprintf("%x", sha1.Sum([]byte(itemByte)))
	return signature == m.Signature
}
