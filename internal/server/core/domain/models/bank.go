package models

import "time"

type BankAccount struct {
	ID             uint64
	CardHolder     string
	Number         uint64
	ExpirationDate time.Time
	SecurityCode   string
	Metadata       string
}

type BankAccountString struct {
	ID             uint64
	CardHolder     string
	Number         uint64
	ExpirationDate string
	SecurityCode   string
	Metadata       string
}
