package handler

type ResponseType int

const (
	OperateOk   ResponseType = 200
	OperateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOk:
		return "Success"
	case OperateFail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}

type RespEntity struct {
	Code  ResponseType
	Msg   string
	Total int
	Data  interface{}
}
