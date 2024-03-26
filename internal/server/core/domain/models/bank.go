package models

import "time"

type BankAccount struct {
	ID               uint64
	CardHolder       string
	Number           uint64
	ExpirationDate   time.Time
	SecurityCodeHash string
	NonceHex         string
	Metadata         string
}

type BankAccountString struct {
	ID               uint64
	CardHolder       string
	Number           uint64
	ExpirationDate   string
	SecurityCodeHash string
	NonceHex         string
	Metadata         string
}
