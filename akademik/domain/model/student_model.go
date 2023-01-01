package model

import (
	"time"
)

type Students struct {
	ID        string    `gorm:"type:varchar(255);primaryKey;column:id"`
	Name      string    `gorm:"type:varchar(255)"`
	Address   string    `gorm:"type:varchar(255)"`
	Gender    string    `gorm:"type:varchar(1)"`
	Email     string    `gorm:"type:varchar(255)"`
	Phone     string    `gorm:"type:varchar(15)"`
	Lecture   []Lecture `gorm:"ForeignKey:StudentID"`
	CreatedAt int64
	UpdatedAt int64
}

func (s *Students) UpdateStudent(student Students) *Students {
	s.Name = student.Name
	s.Address = student.Address
	s.Gender = student.Gender
	s.Email = student.Email
	s.Phone = student.Phone
	s.UpdatedAt = time.Now().UnixMilli()
	return s
}

func (s *Students) UpdateStudentName(name string) *Students {
	s.Name = name
	return s
}
