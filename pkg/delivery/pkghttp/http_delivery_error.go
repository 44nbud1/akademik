package pkghttp

// ErrorHandler represent error handler struct, use strictly for delivery only
type ErrorHandler struct {
	Code     string                 `json:"code,omitempty"`
	Message  string                 `json:"message,omitempty"`
	Detail   string                 `json:"detail,omitempty"`
	Raw      error                  `json:"-"`
	HTTPCode int                    `json:"-"`
	Errors   map[string]interface{} `json:"errors,omitempty"`
}

func (e *ErrorHandler) Error() string {
	return e.Message
}

func (e *ErrorHandler) Unwarp() error {
	return e.Raw
}
