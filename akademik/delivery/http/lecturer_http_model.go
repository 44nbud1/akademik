package http

type Lecturer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (l *Lecturer) setTime(timeNow int64) *Lecturer {

	l.CreatedAt = timeNow
	l.UpdatedAt = timeNow
	return l
}
