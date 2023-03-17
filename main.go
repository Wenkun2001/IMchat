package main

import (
	"IMchat/router"
	"IMchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8089") // 监听并在 localhost:8089 上启动服务
}
