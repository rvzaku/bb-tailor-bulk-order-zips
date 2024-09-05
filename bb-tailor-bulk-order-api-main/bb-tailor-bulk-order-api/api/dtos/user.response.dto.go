package dtos

import "time"

type UserResponseDTO struct {
	ID        string                 `json:"ID"`
	Email     string                 `json:"Email"`
	CreatedAt time.Time              `json:"CreatedAt"`
	UpdatedAt time.Time              `json:"UpdatedAt"`
	Roles     []string               `json:"Roles"`
	Profile   UserProfileResponseDTO `json:"Profile"`
}

type UserProfileResponseDTO struct {
	ID        string `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Phone     string `json:"Phone"`
	Age       int    `json:"Age"`
	Gender    string `json:"Gender"`
}
