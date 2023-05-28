package model

import (
	"time"
)

// Role represents a user role in the system.
type Role string

// Define user roles.
const (
	Candidate  Role = "candidate"
	Intern     Role = "intern"
	HR         Role = "hr"
	Supervisor Role = "supervisor"
)

// User represents a user in the system.
type User struct {
	ID            int64     `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"` // Note: never expose this in JSON/API responses!
	Name          string    `json:"name"`
	LastName      string    `json:"last_name"`
	Role          Role      `json:"role"`
	LoyaltyPoints int       `json:"loyalty_points"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// LoginRequest represents a request to login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegistrationRequest represents a request to register a new user.
type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// UpdateRoleRequest represents a request to update a user's role.
type UpdateRoleRequest struct {
	UserID int64 `json:"user_id"`
	Role   Role  `json:"role"`
}

type UpdateLoyaltyRequest struct {
	LoyaltyPoints int `json:"loyalty_points"`
}
