package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-mock-best-practice/adapter"
	"go-mock-best-practice/entities"
	"time"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository() *MySQLRepository {
	db, err := adapter.Open("mysql", "root:root@tcp(192.168.200.128:3307)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLRepository{
		db: db,
	}
}

func (m *MySQLRepository) CreateStudent(name string) (stu *entities.Student, err error) {
	// 启动事务
	tx, err := m.db.Begin()
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
	result, err := m.db.Exec("insert into students(name) values(?)", name)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	// 2. 然后，给教室1 添加这个学生
	if _, err = m.db.Exec("insert into classroom_1(stu_id) values(?)", id); err != nil {
		return
	}
	stu = &entities.Student{ID: id, Name: name}
	return
}
