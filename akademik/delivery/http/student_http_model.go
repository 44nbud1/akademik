package http

type Students struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (l *Students) setTime(timeNow int64) *Students {

	l.CreatedAt = timeNow
	l.UpdatedAt = timeNow
	return l
}
