package model

type Course struct {
	ID            string    `gorm:"type:varchar(255);primaryKey;column:id"`
	Name          string    `gorm:"type:varchar(255)"`
	CreditsCourse string    `gorm:"type:varchar(100)"`
	Lecture       []Lecture `gorm:"ForeignKey:CourseID"`
	CreatedAt     int64
	UpdatedAt     int64
}

func (c *Course) UpdateCourse(course Course) *Course {
	c.Name = course.Name
	c.CreditsCourse = course.CreditsCourse

	return c
}
