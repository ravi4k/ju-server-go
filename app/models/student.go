package models

type Student struct {
	User       User
	RollNumber string
}

func (Student) TableName() string {
	return "student_info"
}
