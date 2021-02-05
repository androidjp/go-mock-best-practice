package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

var (
	instanceDemoController *DemoController
	initDemoControllerOnce sync.Once
)

type DemoController struct {
}

func GetDemoController() *DemoController {
	initDemoControllerOnce.Do(func() {
		instanceDemoController = &DemoController{}
	})
	return instanceDemoController
}

func (d *DemoController) GetMessage(ctx *gin.Context) {
	fmt.Println(ctx.Request.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", ctx.Request.URL.Path)
	fmt.Println("scheme", ctx.Request.URL.Scheme)
	fmt.Println(ctx.Request.Form["url_long"])
	for k, v := range ctx.Request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	ctx.String(200, "Hello Mike!")
}
