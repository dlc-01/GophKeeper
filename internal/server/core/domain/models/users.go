package models

type UserRole string

const (
	Admin       UserRole = "admin"
	SimilarUser UserRole = "cashier"
)

type User struct {
	ID           *uint64
	Email        string
	Username     string `json:"username"`
	PasswordHash string
}
