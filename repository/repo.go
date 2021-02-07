package repository

import (
	"go-mock-best-practice/adapter"
	"go-mock-best-practice/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository() *MySQLRepository {
	db, err := adapter.Open(mysql.Open("root:root@tcp(192.168.200.128:3307)/test?charset=utf8mb4"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &MySQLRepository{
		db: db,
	}
}

func (m *MySQLRepository) CreateStudent(name string) (stu *entities.Student, err error) {
	// 启动事务
	err = m.db.Transaction(func(tx *gorm.DB) error {
		// 1. 先新增一个学生信息
		newStudent := &entities.Student{
			Name: name,
		}
		if err := tx.Create(newStudent).Error; err != nil {
			return err
		}

		// 2. 然后，给教室1 添加这个学生
		classroomInfo := &entities.Classroom{
			StudentID: newStudent.ID,
		}
		if err = tx.Create(classroomInfo).Error; err != nil {
			return err
		}

		stu = newStudent
		return nil
	})
	return
}

// 查询name中带有相关关键字的N个学生
func (m *MySQLRepository) GetStudents(key string, limit int) ([]*entities.Student, error) {
	var stu []*entities.Student
	err := m.db.Select("name").Where("name LIKE ?", key).Limit(limit).Find(&stu).Error
	return stu, err
}
