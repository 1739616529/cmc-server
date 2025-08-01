package resp

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	error
}

func NewError(code int) Error {
	return Error{Code: code}
}
