package boot

import (
	"goframe-web/app/api/wechat"
	_ "goframe-web/packed"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtimer"
)

func init() {
	interval := 60 * time.Second
	wechat.GetAccessToken()
	gtimer.Add(15*interval, func() {
		g.Log().Info("获取微信access_token")
		wechat.GetAccessToken()
	})
}
