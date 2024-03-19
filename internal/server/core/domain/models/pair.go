package models

type Pair struct {
	ID           uint64
	Username     string
	PasswordHash string
	Metadata     string
}
