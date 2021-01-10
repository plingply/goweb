/*
 * @Author: 彭林
 * @Date: 2021-01-10 15:14:51
 * @LastEditTime: 2021-01-10 15:16:00
 * @LastEditors: 彭林
 * @Description:
 * @FilePath: /goweb/boot/wx-task.go
 */
package boot

import (
	"goframe-web/app/api/wechat"
	"time"

	"github.com/gogf/gf/os/gtimer"
)

func WxTask() {
	interval := 60 * time.Second
	wechat.GetAccessToken()
	gtimer.Add(15*interval, func() {
		wechat.GetAccessToken()
	})
}
