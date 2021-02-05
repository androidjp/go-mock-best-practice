package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mock-best-practice/controller"
	"log"
)

func main() {
	// 设置访问的路由
	r := gin.Default()
	r.GET("/message", controller.GetDemoController().GetMessage)

	// 设置监听的端口
	fmt.Println("Start listening 9090, 尝试请求：http://localhost:9090/message?keyA=valA&url_long=123456")
	if err := r.Run(":9090"); err != nil {
		log.Fatal("ListenAdnServe: ", err)
	}
}
