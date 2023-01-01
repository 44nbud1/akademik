package pkgservice

type Code int

const (
	ErrInternal Code = iota
	ErrBadRequest
	ErrInvalid
	ErrNotFound
)

var codeToDetailMap = map[Code]string{
	ErrInternal:   "internal_error",
	ErrBadRequest: "bad_request",
	ErrInvalid:    "invalid",
	ErrNotFound:   "not_found",
}

func (c Code) String() string {
	return codeToDetailMap[c]
}
