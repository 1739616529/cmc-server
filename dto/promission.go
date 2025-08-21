package dto

type RolePromissionChange struct {
	RoleID     string `json:"roleId" validate:"required"`
	Promission string `json:"promission"`
	Status     int    `json:"status"`
}

type RoleAdd struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description"`
	Promission  string `json:"promission"`
}
