package gomock_service

import gomock_db "go-mock-best-practice/1_gomock/db"

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
