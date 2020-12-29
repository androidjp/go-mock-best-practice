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

func (d *DemoService) InsertData(key, val string) (string, error) {
	var err error
	for i := 0; i < 3; i++ {
		err = d.Repo.Create(key, []byte(val))
		if err == nil {
			break
		}
	}

	if err != nil {
		if err.Error() == "db connection error" {
			return "need to retry", err
		}
		return "", err
	}
	return "success", nil
}

func (d *DemoService) CheckAndUpdateData(key, val string) error {
	data, err := d.Repo.Retrieve(key)
	if data == nil || len(data) == 0 {
		return fmt.Errorf("can not find key %s", key)
	}
	err = d.Repo.Update(key, []byte(val))
	return err
}
