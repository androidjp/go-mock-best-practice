package main

import (
	"fmt"
	"go-mock-best-practice/controller"
	"log"
	"net/http"
)

func main() {
	// 设置访问的路由
	http.HandleFunc("/message", controller.GetDemoController().GetMessage)
	// 设置监听的端口
	fmt.Println("Start listening 9090, 尝试请求：http://localhost:9090/message?keyA=valA&url_long=123456")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAdnServe: ", err)
	}
}
