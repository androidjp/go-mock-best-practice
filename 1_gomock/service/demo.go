package gomock_service

import (
	"fmt"
	gomock_db "go-mock-best-practice/1_gomock/db"
)

// DemoService service业务层对象，手持Repository接口，拥有DB CRUD能力
type DemoService struct {
	Repo gomock_db.Repository
}

// InsertData 插入某条 key-value
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

// CheckAndUpdateData 检索并更新某条记录
func (d *DemoService) CheckAndUpdateData(key, val string) error {
	// 1. 先调用 repo 接口 Retrieve方法
	data, err := d.Repo.Retrieve(key)
	if data == nil || len(data) == 0 {
		return fmt.Errorf("can not find key %s", key)
	}
	// 2. 后调用 repo 接口 Update方法
	err = d.Repo.Update(key, []byte(val))
	return err
}
