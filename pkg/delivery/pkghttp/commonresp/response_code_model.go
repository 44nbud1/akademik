package commonresp

type ResponseCode struct {
	Result   string
	Message  string
	HttpCode int
	CaseCode string
}

func (ec ResponseCode) GetHttp() int {
	return ec.HttpCode
}

func (ec ResponseCode) GetCaseCode() string {
	return ec.CaseCode
}

func (ec ResponseCode) Error() string {
	return ec.Message
}
