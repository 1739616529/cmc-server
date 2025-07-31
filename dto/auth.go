package dto

type CaptchaGet struct {
	Type    string `json:"type" valid:"Required"`
	Account string `json:"account" valid:"Required"`
}
