package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"go-mock-best-practice/adapter"
	"go-mock-best-practice/entities"
	"xorm.io/xorm"
)

type MySQLRepository struct {
	db *xorm.Engine
}

func NewMySQLRepository() *MySQLRepository {
	engine, err := adapter.Open("mysql", "root:root@tcp(192.168.200.128:3307)/test?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	return &MySQLRepository{
		db: engine,
	}
}

func (m *MySQLRepository) CreateStudent(name string) (stu *entities.Student, err error) {
	// 启动事务
	_, err = m.db.Transaction(func(session *xorm.Session) (interface{}, error) {

		// 1. 先新增一个学生信息
		newStudent := &entities.Student{
			Name: name,
		}
		if _, err := session.InsertOne(newStudent); err != nil {
			return nil, err
		}

		// 2. 然后，给教室1 添加这个学生
		classroomInfo := &entities.Classroom{
			StudentID: newStudent.ID,
		}
		if _, err = session.InsertOne(classroomInfo); err != nil {
			return nil, err
		}
		stu = newStudent
		return nil, nil
	})
	return
}

// 查询name中带有相关关键字的N个学生
func (m *MySQLRepository) GetStudents(key string, limit int) ([]*entities.Student, error) {
	var stu []*entities.Student
	err := m.db.Select("name").Where("name LIKE ?", key).Limit(limit).Find(&stu)
	return stu, err
}
