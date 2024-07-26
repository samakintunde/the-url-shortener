// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
	"time"
)

type EmailVerification struct {
	ID         int64
	UserID     string
	Email      string
	Code       string
	CreatedAt  time.Time
	ExpiresAt  time.Time
	VerifiedAt sql.NullTime
}

type User struct {
	ID          string
	Email       string
	FirstName   sql.NullString
	LastName    sql.NullString
	Password    string
	Status      interface{}
	LastLoginAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
