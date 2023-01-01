package model

import (
	"time"
)

type Lecturer struct {
	ID        string    `gorm:"type:varchar(255);primaryKey;column:id"`
	Name      string    `gorm:"type:varchar(255)"`
	Lecture   []Lecture `gorm:"ForeignKey:LecturerID"`
	CreatedAt int64
	UpdatedAt int64
}

func (l *Lecturer) UpdateLecturer(lecture Lecturer) *Lecturer {

	l.Name = lecture.Name
	l.UpdatedAt = time.Now().UnixMilli()
	return l
}
