package dto

type UserCreate struct {
	Name   string `json:"username" valid:"MaxSize(20)"`
	Phone  string `json:"phone"`
	Passwd string `json:"password"`
	Email  string `json:"email" valid:"Email"`
}

type UserOutput struct {
	Name  string `json:"username" valid:"MaxSize(20)"`
	Phone string `json:"phone" `
	Email string `json:"email" valid:"Email"`
}

type UserLogin struct {
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Passwd string `json:"password" valid:"Required"`
}

type UserRegister struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Passwd   string `json:"password" valid:"Required"`
	Type     string `json:"type" valid:"Required"`
	Name     string `json:"name" valid:"Required"`
	Code     string `json:"code" valid:"Required"`
	VerifyId string `json:"verifyId" valid:"Required"`
}
