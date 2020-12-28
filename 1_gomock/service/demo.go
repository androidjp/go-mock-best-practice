package gomock_service

import (
	"fmt"
	gomock_db "go-mock-best-practice/1_gomock/db"
)

type DemoService struct {
	Repo gomock_db.Repository
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
