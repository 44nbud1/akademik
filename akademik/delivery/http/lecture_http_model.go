package http

type Lecture struct {
	ID         string `json:"id,omitempty"`
	StudentID  string `json:"studentId,omitempty"`
	LecturerID string `json:"lecturerId,omitempty"`
	Score      string `json:"score,omitempty"`
	Name       string `json:"name,omitempty"`
	CourseID   string `json:"courseId,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty"`
	UpdatedAt  int64  `json:"updatedAt,omitempty"`
}

func (l *Lecture) setTime(timeNow int64) *Lecture {

	l.CreatedAt = timeNow
	l.UpdatedAt = timeNow
	return l
}
