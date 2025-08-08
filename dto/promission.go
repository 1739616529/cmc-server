package dto

type RolePromissionChange struct {
	RoleID     string `json:"roleId" validate:"required"`
	Promission string `json:"promission"`
	Status     int    `json:"status"`
}
