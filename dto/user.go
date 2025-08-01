package dto

type UserCreate struct {
	Name   string `json:"username" valid:"MaxSize(20)"`
	Phone  string `json:"phone"`
	Passwd string `json:"passwd"`
	Email  string `json:"email" valid:"Email"`
}

type UserOutput struct {
	Id     string `json:"id"`
	Name   string `json:"username" `
	Phone  string `json:"phone" `
	Email  string `json:"email" `
	Avatar string `json:"avatar" `
}

type UserLogin struct {
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Passwd string `json:"passwd" valid:"Required"`
}

type UserRegister struct {
	Account  string `json:"account" valid:"Required"`
	Type     string `json:"type" valid:"Required"`
	Code     string `json:"code" valid:"Required"`
	VerifyId string `json:"verifyId" valid:"Required"`
}
