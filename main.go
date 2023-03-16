package main

import (
	"IMchat/router"
)

func main() {
	r := router.Router()
	r.Run(":8089") // 监听并在 0.0.0.0:8089 上启动服务
}
