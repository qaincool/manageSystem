package request

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
