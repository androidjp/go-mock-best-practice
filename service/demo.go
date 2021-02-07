package service

import (
	"database/sql"
	"errors"
	"fmt"
	"go-mock-best-practice/entities"
	"go-mock-best-practice/repository"
)

type DemoService struct {
}

// 情况一：直接将db对象作为入参
func (d *DemoService) AddStudentDirectly(db *sql.DB, name string) (stu *entities.Student, err error) {
	// 启动事务
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 1. 先新增一个学生信息
	result, err := db.Exec("insert into students(name) values(?)", name)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	// 2. 然后，给教室1 添加这个学生
	if _, err = db.Exec("insert into classroom_1(stu_id) values(?)", id); err != nil {
		return
	}
	stu = &entities.Student{ID: id, Name: name}
	return
}

// 情况二：有自己的repository层对象，去封装db的相关操作
func (d *DemoService) AddStudentByName(name string) error {
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	student, err := repository.NewMySQLRepository().CreateStudent(name)
	if err != nil {
		return err
	}
	fmt.Println(student.ID)
	fmt.Println(student.Name)
	return nil
}
