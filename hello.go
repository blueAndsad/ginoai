package main

import (
	"awesomeProject/routers"
)

func main() {
	//注册路由
	r := routers.SetupRouter()
	r.Run(":8080") //开启服务器，默认监听localhost:8080
}
