package pkgservice

type ErrorService struct {
	Code      Code
	Message   string
	Detail    string
	Raw       error
	Attribute map[string]string
}

func (e *ErrorService) GetHttp() int {
	return int(e.Code)
}

func (e *ErrorService) GetCaseCode() string {
	return e.Code.String()
}

func (e *ErrorService) GetCode() Code {
	return e.Code
}

func (e *ErrorService) String() string {
	return e.Message
}

func (e *ErrorService) Error() string {
	return e.Message
}

func NewErrorService(code Code) *ErrorService {
	return &ErrorService{
		Code:    code,
		Message: code.String(),
	}
}
