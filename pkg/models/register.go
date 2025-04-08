package models

type RegisterRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	InviteCode string `json:"invite_code"`
}
