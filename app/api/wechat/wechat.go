/*
 * @Author: 彭林
 * @Date: 2021-01-10 11:56:43
 * @LastEditTime: 2021-01-10 14:55:43
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/app/api/wechat/wechat.go
 */
package wechat

import (
	"crypto/sha1"
	"fmt"
	"goframe-web/library/cache"
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

	result := data.CheckSignature(token)
	if result {
		r.Response.WriteExit(data.EchoStr)
	} else {
		r.Response.WriteExit(false)
	}
}

/**
 * @description: 获取微信access_token
 * @param {*ghttp.Request} r
 * @return {*}
 */
func GetAccessToken() {
	appid := g.Cfg().GetString("wechat.appid")
	appsecret := g.Cfg().GetString("wechat.appsecret")
	if r, err := g.Client().Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + appsecret); err != nil {
		panic(err)
	} else {
		defer r.Close()
		cache.Set("wechat_access_token", r.ReadAllString(), 0)
	}
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
