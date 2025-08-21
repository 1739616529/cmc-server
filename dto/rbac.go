package dto

import (
	"time"
)

type RoleOutput struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	IsBuiltIn   bool      `json:"isBuiltIn"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PermissionOutput struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Bit         string    `json:"bit"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
