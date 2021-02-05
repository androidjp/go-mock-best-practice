package gomock_service

import (
	"fmt"
	gomock_db "go-mock-best-practice/1_gomock/db"
)

var (
	GlobalCount int
	Host        string
)

type DemoService struct {
	Count int
	Repo  gomock_db.Repository
}

// 输出此时的全局变量 和 成员变量
// 方法
func (d *DemoService) CheckConnect() string {
	return fmt.Sprintf("%d:%s:%d", GlobalCount, Host, d.Count)
}

// 函数
var Exec = func(cmd string, args ...string) (string, error) {
	return "", nil
}

// 过程
var DestroyResource = func() {
	fmt.Println("清理资源等工作")
}
