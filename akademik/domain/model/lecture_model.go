package model

type Lecture struct {
	ID         string `gorm:"type:varchar(255);primaryKey;column:id"`
	Name       string
	StudentID  string `gorm:"ForeignKey:ID"`
	LecturerID string `gorm:"ForeignKey:ID"`
	CourseID   string `gorm:"ForeignKey:ID"`
	Score      string `gorm:"type:varchar(3)"`
	CreatedAt  int64
	UpdatedAt  int64
}

func (l *Lecture) UpdateLecture(lecture Lecture) *Lecture {

	l.Score = lecture.Score
	return l
}

func (l *Lecture) UpdateLectureName(lecture Lecture) *Lecture {

	l.Name = lecture.Name
	return l
}
