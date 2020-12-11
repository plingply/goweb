package main

import (
	_ "goframe-web/boot"
	_ "goframe-web/router"
	"github.com/gogf/gf/frame/g"
)

func main() {

	//启动服务
	g.Server().Run()
}
