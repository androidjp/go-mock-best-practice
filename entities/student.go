package entities

type Student struct {
	ID   uint   `json:"id" xorm:"pk id autoincr"`
	Name string `json:"name" xorm:"'name'"`
}

func (Student) TableName() string {
	return "students"
}

type Classroom struct {
	ID        uint `json:"id" xorm:"pk id autoincr"`
	StudentID uint `json:"stu_id" xorm:"'stu_id'"`
}

func (Classroom) TableName() string {
	return "classroom_1"
}
