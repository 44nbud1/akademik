package http

type Course struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	CreditsCourse string `json:"creditsCourse,omitempty"`
	CreatedAt     int64  `json:"createdAt,omitempty"`
	UpdatedAt     int64  `json:"updatedAt,omitempty"`
}

func (l *Course) setTime(timeNow int64) *Course {

	l.CreatedAt = timeNow
	l.UpdatedAt = timeNow
	return l
}
