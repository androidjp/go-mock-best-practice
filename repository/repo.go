package repository

import "go-mock-best-practice/entities"

type Repository interface {
	// 检索
	Retrieve(key string, movie *entities.Movie) error
}

func GetInstance() Repository {
	return nil
}
