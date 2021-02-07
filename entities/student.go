package entities

type Student struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

type Classroom struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	StudentID uint `json:"stu_id" gorm:"column:stu_id"`
}

func (Classroom) TableName() string {
	return "classroom_1"
}
