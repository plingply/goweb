/*
 * @Author: 彭林
 * @Date: 2021-01-10 11:56:43
 * @LastEditTime: 2021-01-11 14:16:24
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
	var token *wechatAccessToken
	if err := g.Client().GetVar("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + appsecret).Scan(&token); err != nil {
		panic(err)
	} else {
		cache.Set("wechat_access_token", token.AccessToken, 0)
		g.Log().Info("获取微信access_token:" + token.AccessToken)
	}
}

type messageEntity struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

type wechatAccessToken struct {
	AccessToken string `json:"access_token"`
}

func (m *messageEntity) CheckSignature(token string) bool {
	item := []string{token, m.Timestamp, m.Nonce}
	sort.Strings(item)
	itemByte := strings.Join(item, "")
	signature := fmt.Sprintf("%x", sha1.Sum([]byte(itemByte)))
	return signature == m.Signature
}
