package entity

import "go-mock-best-practice/1_gomock/adapter"

type Student struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func (s *Student) Print() (string, error) {
	bytes, err := adapter.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
