package controller

import (
	"fmt"
	"net/http"
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

func (d *DemoController) GetMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Mike!") // 这个写入到 w 的是输出到客户端的
}
